import { z } from 'zod';

import { makeAPI } from './makeAPI';
import { MigrationSchema } from './models/migration';

export const ResponseSchema = z.object({
  migration: MigrationSchema,
});

export type Response = z.infer<typeof ResponseSchema>;

export const makeGetMigrationAPI = (cosmosAddress: string) =>
  makeAPI({
    url: `/likenft/migration/${cosmosAddress}`,
    method: 'GET',
    responseSchema: ResponseSchema,
  });
