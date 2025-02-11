import { z } from 'zod';

export const MigrationClassStatusSchema = z.enum([
  'pending',
  'scheduled',
  'in_progress',
  'success',
  'failed',
]);

export const MigrationClassSchema = z.object({
  id: z.string(),
  name: z.string(),
  likerland_url: z.string(),
  image: z.string(),
  status: MigrationClassStatusSchema,
  enqueued_time: z.string().nullable(),
  evm_tx_hash: z.string().nullable(),
  failed_reason: z.string().nullable(),
});

export const MigrationNFTStatusSchema = z.enum([
  'pending',
  'scheduled',
  'in_progress',
  'success',
  'failed',
]);

export const MigrationNFTSchema = z.object({
  class_id: z.string(),
  id: z.string(),
  name: z.string(),
  likerland_url: z.string(),
  image: z.string(),
  status: MigrationNFTStatusSchema,
  enqueued_time: z.string().nullable(),
  evm_tx_hash: z.string().nullable(),
  failed_reason: z.string().nullable(),
});

export const MigrationStatusSchema = z.enum([
  'preview',
  'pending',
  'in_progress',
  'success',
  'failed',
]);

export const MigrationSchema = z.object({
  status: MigrationStatusSchema,
  classes: z.array(MigrationClassSchema),
  nfts: z.array(MigrationNFTSchema),
});

export type Migration = z.infer<typeof MigrationSchema>;
