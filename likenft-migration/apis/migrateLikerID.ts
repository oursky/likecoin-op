import { z } from 'zod';

import { makeAPI } from './makeAPI';

export const RequestSchema = z.object({
  cosmos_address: z.string(),
  cosmos_pub_key: z.string(),
  like_id: z.string().nullable(),
  eth_address: z.string(),
  cosmos_signature: z.string(),
  cosmos_signing_message: z.string(),
  eth_signature: z.string(),
  eth_signing_message: z.string(),
});

export const ResponseSchema = z.object({
  message: z.string(),
});

export const makeMigrateLikerIDAPI = makeAPI({
  url: '/likenft/likerid/migration',
  method: 'POST',
  requestSchema: RequestSchema,
  responseSchema: ResponseSchema,
});
