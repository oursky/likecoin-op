package database

import (
	"context"
	"time"

	"likenft-indexer/ent"
	"likenft-indexer/ent/nftclass"

	"entgo.io/ent/dialect/sql"
)

type NFTClassAcquireBookNFTEventsRepository interface {
	RetrieveState(
		ctx context.Context,
		contractAddress string,
	) (*ent.NFTClass, error)

	RequestForEnqueue(
		ctx context.Context,
		count int,
	) ([]*ent.NFTClass, error)

	Enqueued(
		ctx context.Context,
		m *ent.NFTClass,
	) error

	// Should reset the state
	// and to be enqueued in next batch
	EnqueueFailed(
		ctx context.Context,
		m *ent.NFTClass,
		err error,
	) error

	Processing(
		ctx context.Context,
		m *ent.NFTClass,
	) error

	// Should reset the state
	// and update attributes
	Completed(
		ctx context.Context,
		m *ent.NFTClass,
		lastProcessedTime time.Time,
		newScore float64,
	) error

	// Should reset the state
	// and to be enqueued in next batch
	Failed(
		ctx context.Context,
		m *ent.NFTClass,
		err error,
	) error
}

type nftClassAcquireBookNFTEventsRepository struct {
	dbService Service
}

func MakeNFTClassAcquireBookNFTEventsRepository(
	dbService Service,
) NFTClassAcquireBookNFTEventsRepository {
	return &nftClassAcquireBookNFTEventsRepository{
		dbService: dbService,
	}
}

func (r *nftClassAcquireBookNFTEventsRepository) RetrieveState(
	ctx context.Context,
	contractAddress string,
) (*ent.NFTClass, error) {
	return r.dbService.Client().NFTClass.Query().Where(
		nftclass.AddressEqualFold(contractAddress),
	).Only(ctx)
}

func (r *nftClassAcquireBookNFTEventsRepository) RequestForEnqueue(
	ctx context.Context,
	count int,
) ([]*ent.NFTClass, error) {
	resChan := make(chan []*ent.NFTClass, 1)

	err := WithTx(ctx, r.dbService.Client(), func(tx *ent.Tx) error {
		var addressRows []struct {
			Address string `json:"address"`
		}
		err := tx.NFTClass.Query().
			Select(nftclass.FieldAddress).
			Where(
				nftclass.DisabledForIndexingEQ(false),
				nftclass.Or(
					nftclass.AcquireBookNftEventsStatusIsNil(),
					nftclass.AcquireBookNftEventsStatusIn(
						nftclass.AcquireBookNftEventsStatusCompleted,
						nftclass.AcquireBookNftEventsStatusEnqueueFailed,
						nftclass.AcquireBookNftEventsStatusFailed,
					),
				),
			).Order(
			nftclass.ByAcquireBookNftEventsScore(sql.OrderNullsFirst(), sql.OrderAsc()),
		).Limit(count).
			Select(nftclass.FieldAddress).
			Scan(ctx, &addressRows)

		if err != nil {
			return err
		}

		if len(addressRows) == 0 {
			resChan <- []*ent.NFTClass{}
			return nil
		}

		addresses := make([]string, len(addressRows))
		for i, addressRow := range addressRows {
			addresses[i] = addressRow.Address
		}

		if err := tx.NFTClass.Update().
			SetAcquireBookNftEventsStatus(nftclass.AcquireBookNftEventsStatusEnqueueing).
			Where(
				nftclass.AddressIn(addresses...),
			).
			Exec(ctx); err != nil {
			return err
		}

		nftClasses, err := tx.NFTClass.Query().
			Where(
				nftclass.AddressIn(addresses...),
			).Order(
			nftclass.ByAcquireBookNftEventsScore(sql.OrderNullsFirst(), sql.OrderAsc()),
		).All(ctx)

		resChan <- nftClasses
		return nil
	})

	if err != nil {
		return nil, err
	}
	return <-resChan, nil
}

func (r *nftClassAcquireBookNFTEventsRepository) Enqueued(
	ctx context.Context,
	m *ent.NFTClass,
) error {
	return WithTx(ctx, r.dbService.Client(), func(tx *ent.Tx) error {
		return tx.NFTClass.UpdateOne(m).
			Where(
				nftclass.AddressEqualFold(m.Address),
			).
			SetAcquireBookNftEventsStatus(nftclass.AcquireBookNftEventsStatusEnqueued).
			Exec(ctx)
	})
}

func (r *nftClassAcquireBookNFTEventsRepository) EnqueueFailed(
	ctx context.Context,
	m *ent.NFTClass,
	err error,
) error {

	return WithTx(ctx, r.dbService.Client(), func(tx *ent.Tx) error {
		return tx.NFTClass.UpdateOne(m).
			Where(
				nftclass.AddressEqualFold(m.Address),
			).
			SetAcquireBookNftEventsStatus(
				nftclass.AcquireBookNftEventsStatusEnqueueFailed,
			).
			SetAcquireBookNftEventsFailedReason(
				err.Error(),
			).
			Exec(ctx)
	})
}

func (r *nftClassAcquireBookNFTEventsRepository) Processing(
	ctx context.Context,
	m *ent.NFTClass,
) error {
	return WithTx(ctx, r.dbService.Client(), func(tx *ent.Tx) error {
		return tx.NFTClass.UpdateOne(m).
			Where(
				nftclass.AddressEqualFold(m.Address),
			).
			SetAcquireBookNftEventsStatus(nftclass.AcquireBookNftEventsStatusProcessing).
			Exec(ctx)
	})
}

func (r *nftClassAcquireBookNFTEventsRepository) Completed(
	ctx context.Context,
	m *ent.NFTClass,
	lastProcessedTime time.Time,
	newScore float64,
) error {
	return WithTx(ctx, r.dbService.Client(), func(tx *ent.Tx) error {
		return tx.NFTClass.UpdateOne(m).
			Where(
				nftclass.AddressEqualFold(m.Address),
			).
			SetAcquireBookNftEventsStatus(nftclass.AcquireBookNftEventsStatusCompleted).
			SetAcquireBookNftEventsLastProcessedTime(lastProcessedTime).
			SetAcquireBookNftEventsScore(newScore).
			Exec(ctx)
	})
}

func (r *nftClassAcquireBookNFTEventsRepository) Failed(
	ctx context.Context,
	m *ent.NFTClass,
	err error,
) error {
	return WithTx(ctx, r.dbService.Client(), func(tx *ent.Tx) error {
		return tx.NFTClass.UpdateOne(m).
			Where(
				nftclass.AddressEqualFold(m.Address),
			).
			SetAcquireBookNftEventsStatus(nftclass.AcquireBookNftEventsStatusFailed).
			SetAcquireBookNftEventsFailedReason(err.Error()).
			Exec(ctx)
	})
}
