import { z } from 'zod';

import { makeAPI } from './makeAPI';
import { MigrationSchema } from './models/migration';

export const RequestSchema = z.object({
  cosmos_address: z.string(),
  eth_address: z.string(),
});

export type Request = z.infer<typeof RequestSchema>;

export const ResponseSchema = z.object({
  migration: MigrationSchema,
});

export type Response = z.infer<typeof ResponseSchema>;

export const makeCreateMigrationAPI = makeAPI({
  url: `/likenft/migration`,
  method: 'POST',
  requestSchema: RequestSchema,
  responseSchema: ResponseSchema,
});
