import { z } from "zod";

import { makeAPI } from "./makeAPI";
import { LikeNFTMigrationDetailSchema } from "./models/likenftMigration";

export const RequestSchema = z.object({
  asset_snapshot_id: z.number(),
  cosmos_address: z.string(),
  eth_address: z.string(),
});

export type Request = z.infer<typeof RequestSchema>;

export const ResponseSchema = z.object({
  migration: LikeNFTMigrationDetailSchema,
});

export type Response = z.infer<typeof ResponseSchema>;

export const makeCreateMigrationAPI = makeAPI({
  url: `/likenft/migration`,
  method: "POST",
  requestSchema: RequestSchema,
  responseSchema: ResponseSchema,
});
