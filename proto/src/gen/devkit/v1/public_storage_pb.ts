// @generated by protoc-gen-es v2.2.2 with parameter "target=ts"
// @generated from file devkit/v1/public_storage.proto (package devkit.v1, syntax proto3)
/* eslint-disable */

import type { GenFile, GenMessage } from "@bufbuild/protobuf/codegenv1";
import { fileDesc, messageDesc } from "@bufbuild/protobuf/codegenv1";
import { file_google_protobuf_struct } from "@bufbuild/protobuf/wkt";
import { file_buf_validate_validate } from "../../buf/validate/validate_pb";
import type { JsonObject, Message } from "@bufbuild/protobuf";

/**
 * Describes the file devkit/v1/public_storage.proto.
 */
export const file_devkit_v1_public_storage: GenFile = /*@__PURE__*/
  fileDesc("Ch5kZXZraXQvdjEvcHVibGljX3N0b3JhZ2UucHJvdG8SCWRldmtpdC52MSJrChFGaWxlQ3JlYXRlUmVxdWVzdBIVCgRwYXRoGAEgASgJQge6SARyAhACEhwKC2J1Y2tldF9uYW1lGAIgASgJQge6SARyAhACEg4KBnJlYWRlchgDIAEoDBIRCglmaWxlX3R5cGUYBCABKAkiRAoVRmlsZUNyZWF0ZUJ1bGtSZXF1ZXN0EisKBWZpbGVzGAEgAygLMhwuZGV2a2l0LnYxLkZpbGVDcmVhdGVSZXF1ZXN0IiIKEkZpbGVDcmVhdGVSZXNwb25zZRIMCgRwYXRoGAEgASgJIiYKFkZpbGVDcmVhdGVCdWxrUmVzcG9uc2USDAoEcGF0aBgBIAMoCSJqChJJbXBvcnRUYWJsZVJlcXVlc3QSGwoKdGFibGVfbmFtZRgBIAEoCUIHukgEcgIQAhITCgtzY2hlbWFfbmFtZRgCIAEoCRISCgpzaGVldF9uYW1lGAMgASgJEg4KBnJlYWRlchgEIAEoDCImChNJbXBvcnRUYWJsZVJlc3BvbnNlEg8KB21lc3NhZ2UYASABKAkipQEKDVN0b3JhZ2VCdWNrZXQSCgoCaWQYASABKAkSDAoEbmFtZRgCIAEoCRIOCgZwdWJsaWMYAyABKAgSEgoKY3JlYXRlZF9hdBgEIAEoCRISCgp1cGRhdGVkX2F0GAUgASgJEg0KBW93bmVyGAYgASgJEhcKD2ZpbGVfc2l6ZV9saW1pdBgHIAEoAxIaChJhbGxvd2VkX21pbWVfdHlwZXMYCCADKAkimwEKGUJ1Y2tldENyZWF0ZVVwZGF0ZVJlcXVlc3QSEwoLYnVja2V0X25hbWUYASABKAkSEAoIaXNfcHVsaWMYAiABKAgSKAoPZmlsZV9zaXplX2xpbWl0GAMgASgJQg+6SAxyCjIIXlswLTldKyQSGgoSYWxsb3dlZF9maWxlX3R5cGVzGAQgAygJEhEKCWlzX3VwZGF0ZRgFIAEoCCJGChpCdWNrZXRDcmVhdGVVcGRhdGVSZXNwb25zZRIoCgZidWNrZXQYASABKAsyGC5kZXZraXQudjEuU3RvcmFnZUJ1Y2tldCITChFCdWNrZXRMaXN0UmVxdWVzdCI/ChJCdWNrZXRMaXN0UmVzcG9uc2USKQoHYnVja2V0cxgBIAMoCzIYLmRldmtpdC52MS5TdG9yYWdlQnVja2V0IuABCgpGaWxlT2JqZWN0EgwKBG5hbWUYASABKAkSEQoJYnVja2V0X2lkGAIgASgJEg0KBW93bmVyGAMgASgJEgoKAmlkGAQgASgJEhIKCnVwZGF0ZWRfYXQYBSABKAkSEgoKY3JlYXRlZF9hdBgGIAEoCRIYChBsYXN0X2FjY2Vzc2VkX2F0GAcgASgJEikKCG1ldGFkYXRhGAggASgLMhcuZ29vZ2xlLnByb3RvYnVmLlN0cnVjdBIpCgdidWNrZXRzGAkgASgLMhguZGV2a2l0LnYxLlN0b3JhZ2VCdWNrZXQiVwoPRmlsZUxpc3RSZXF1ZXN0EhEKCWJ1Y2tldF9pZBgBIAEoCRISCgpxdWVyeV9wYXRoGAIgASgJEg0KBWxpbWl0GAMgASgFEg4KBm9mZmVzdBgEIAEoBSI4ChBGaWxlTGlzdFJlc3BvbnNlEiQKBWZpbGVzGAEgAygLMhUuZGV2a2l0LnYxLkZpbGVPYmplY3QiUQoRRmlsZURlbGV0ZVJlcXVlc3QSGgoJYnVja2V0X2lkGAEgASgJQge6SARyAhACEiAKC2ZpbGVzX3BhdGhzGAIgAygJQgu6SAiSAQUQ9AMYASJGChJGaWxlRGVsZXRlUmVzcG9uc2USMAoJcmVzcG9uc2VzGAEgAygLMh0uZGV2a2l0LnYxLkZpbGVDcmVhdGVSZXNwb25zZUKnAQoNY29tLmRldmtpdC52MUISUHVibGljU3RvcmFnZVByb3RvUAFaPWdpdGh1Yi5jb20vZGFyd2lzaGRldi9kZXZraXQtYXBpL3Byb3RvX2dlbi9kZXZraXQvdjE7ZGV2a2l0djGiAgNEWFiqAglEZXZraXQuVjHKAglEZXZraXRcVjHiAhVEZXZraXRcVjFcR1BCTWV0YWRhdGHqAgpEZXZraXQ6OlYxYgZwcm90bzM", [file_google_protobuf_struct, file_buf_validate_validate]);

/**
 * @generated from message devkit.v1.FileCreateRequest
 */
export type FileCreateRequest = Message<"devkit.v1.FileCreateRequest"> & {
  /**
   * @generated from field: string path = 1;
   */
  path: string;

  /**
   * @generated from field: string bucket_name = 2;
   */
  bucketName: string;

  /**
   * @generated from field: bytes reader = 3;
   */
  reader: Uint8Array;

  /**
   * @generated from field: string file_type = 4;
   */
  fileType: string;
};

/**
 * Describes the message devkit.v1.FileCreateRequest.
 * Use `create(FileCreateRequestSchema)` to create a new message.
 */
export const FileCreateRequestSchema: GenMessage<FileCreateRequest> = /*@__PURE__*/
  messageDesc(file_devkit_v1_public_storage, 0);

/**
 * @generated from message devkit.v1.FileCreateBulkRequest
 */
export type FileCreateBulkRequest = Message<"devkit.v1.FileCreateBulkRequest"> & {
  /**
   * @generated from field: repeated devkit.v1.FileCreateRequest files = 1;
   */
  files: FileCreateRequest[];
};

/**
 * Describes the message devkit.v1.FileCreateBulkRequest.
 * Use `create(FileCreateBulkRequestSchema)` to create a new message.
 */
export const FileCreateBulkRequestSchema: GenMessage<FileCreateBulkRequest> = /*@__PURE__*/
  messageDesc(file_devkit_v1_public_storage, 1);

/**
 * @generated from message devkit.v1.FileCreateResponse
 */
export type FileCreateResponse = Message<"devkit.v1.FileCreateResponse"> & {
  /**
   * @generated from field: string path = 1;
   */
  path: string;
};

/**
 * Describes the message devkit.v1.FileCreateResponse.
 * Use `create(FileCreateResponseSchema)` to create a new message.
 */
export const FileCreateResponseSchema: GenMessage<FileCreateResponse> = /*@__PURE__*/
  messageDesc(file_devkit_v1_public_storage, 2);

/**
 * @generated from message devkit.v1.FileCreateBulkResponse
 */
export type FileCreateBulkResponse = Message<"devkit.v1.FileCreateBulkResponse"> & {
  /**
   * @generated from field: repeated string path = 1;
   */
  path: string[];
};

/**
 * Describes the message devkit.v1.FileCreateBulkResponse.
 * Use `create(FileCreateBulkResponseSchema)` to create a new message.
 */
export const FileCreateBulkResponseSchema: GenMessage<FileCreateBulkResponse> = /*@__PURE__*/
  messageDesc(file_devkit_v1_public_storage, 3);

/**
 * @generated from message devkit.v1.ImportTableRequest
 */
export type ImportTableRequest = Message<"devkit.v1.ImportTableRequest"> & {
  /**
   * @generated from field: string table_name = 1;
   */
  tableName: string;

  /**
   * @generated from field: string schema_name = 2;
   */
  schemaName: string;

  /**
   * @generated from field: string sheet_name = 3;
   */
  sheetName: string;

  /**
   * @generated from field: bytes reader = 4;
   */
  reader: Uint8Array;
};

/**
 * Describes the message devkit.v1.ImportTableRequest.
 * Use `create(ImportTableRequestSchema)` to create a new message.
 */
export const ImportTableRequestSchema: GenMessage<ImportTableRequest> = /*@__PURE__*/
  messageDesc(file_devkit_v1_public_storage, 4);

/**
 * @generated from message devkit.v1.ImportTableResponse
 */
export type ImportTableResponse = Message<"devkit.v1.ImportTableResponse"> & {
  /**
   * @generated from field: string message = 1;
   */
  message: string;
};

/**
 * Describes the message devkit.v1.ImportTableResponse.
 * Use `create(ImportTableResponseSchema)` to create a new message.
 */
export const ImportTableResponseSchema: GenMessage<ImportTableResponse> = /*@__PURE__*/
  messageDesc(file_devkit_v1_public_storage, 5);

/**
 * @generated from message devkit.v1.StorageBucket
 */
export type StorageBucket = Message<"devkit.v1.StorageBucket"> & {
  /**
   * @generated from field: string id = 1;
   */
  id: string;

  /**
   * @generated from field: string name = 2;
   */
  name: string;

  /**
   * @generated from field: bool public = 3;
   */
  public: boolean;

  /**
   * @generated from field: string created_at = 4;
   */
  createdAt: string;

  /**
   * @generated from field: string updated_at = 5;
   */
  updatedAt: string;

  /**
   * Optional fields
   *
   * @generated from field: string owner = 6;
   */
  owner: string;

  /**
   * @generated from field: int64 file_size_limit = 7;
   */
  fileSizeLimit: bigint;

  /**
   * @generated from field: repeated string allowed_mime_types = 8;
   */
  allowedMimeTypes: string[];
};

/**
 * Describes the message devkit.v1.StorageBucket.
 * Use `create(StorageBucketSchema)` to create a new message.
 */
export const StorageBucketSchema: GenMessage<StorageBucket> = /*@__PURE__*/
  messageDesc(file_devkit_v1_public_storage, 6);

/**
 * @generated from message devkit.v1.BucketCreateUpdateRequest
 */
export type BucketCreateUpdateRequest = Message<"devkit.v1.BucketCreateUpdateRequest"> & {
  /**
   * @generated from field: string bucket_name = 1;
   */
  bucketName: string;

  /**
   * @generated from field: bool is_pulic = 2;
   */
  isPulic: boolean;

  /**
   * @generated from field: string file_size_limit = 3;
   */
  fileSizeLimit: string;

  /**
   * @generated from field: repeated string allowed_file_types = 4;
   */
  allowedFileTypes: string[];

  /**
   * @generated from field: bool is_update = 5;
   */
  isUpdate: boolean;
};

/**
 * Describes the message devkit.v1.BucketCreateUpdateRequest.
 * Use `create(BucketCreateUpdateRequestSchema)` to create a new message.
 */
export const BucketCreateUpdateRequestSchema: GenMessage<BucketCreateUpdateRequest> = /*@__PURE__*/
  messageDesc(file_devkit_v1_public_storage, 7);

/**
 * @generated from message devkit.v1.BucketCreateUpdateResponse
 */
export type BucketCreateUpdateResponse = Message<"devkit.v1.BucketCreateUpdateResponse"> & {
  /**
   * @generated from field: devkit.v1.StorageBucket bucket = 1;
   */
  bucket?: StorageBucket;
};

/**
 * Describes the message devkit.v1.BucketCreateUpdateResponse.
 * Use `create(BucketCreateUpdateResponseSchema)` to create a new message.
 */
export const BucketCreateUpdateResponseSchema: GenMessage<BucketCreateUpdateResponse> = /*@__PURE__*/
  messageDesc(file_devkit_v1_public_storage, 8);

/**
 * @generated from message devkit.v1.BucketListRequest
 */
export type BucketListRequest = Message<"devkit.v1.BucketListRequest"> & {
};

/**
 * Describes the message devkit.v1.BucketListRequest.
 * Use `create(BucketListRequestSchema)` to create a new message.
 */
export const BucketListRequestSchema: GenMessage<BucketListRequest> = /*@__PURE__*/
  messageDesc(file_devkit_v1_public_storage, 9);

/**
 * @generated from message devkit.v1.BucketListResponse
 */
export type BucketListResponse = Message<"devkit.v1.BucketListResponse"> & {
  /**
   * @generated from field: repeated devkit.v1.StorageBucket buckets = 1;
   */
  buckets: StorageBucket[];
};

/**
 * Describes the message devkit.v1.BucketListResponse.
 * Use `create(BucketListResponseSchema)` to create a new message.
 */
export const BucketListResponseSchema: GenMessage<BucketListResponse> = /*@__PURE__*/
  messageDesc(file_devkit_v1_public_storage, 10);

/**
 * @generated from message devkit.v1.FileObject
 */
export type FileObject = Message<"devkit.v1.FileObject"> & {
  /**
   * @generated from field: string name = 1;
   */
  name: string;

  /**
   * @generated from field: string bucket_id = 2;
   */
  bucketId: string;

  /**
   * @generated from field: string owner = 3;
   */
  owner: string;

  /**
   * @generated from field: string id = 4;
   */
  id: string;

  /**
   * @generated from field: string updated_at = 5;
   */
  updatedAt: string;

  /**
   * @generated from field: string created_at = 6;
   */
  createdAt: string;

  /**
   * @generated from field: string last_accessed_at = 7;
   */
  lastAccessedAt: string;

  /**
   * Use google.protobuf.Struct for dynamic metadata
   *
   * @generated from field: google.protobuf.Struct metadata = 8;
   */
  metadata?: JsonObject;

  /**
   * @generated from field: devkit.v1.StorageBucket buckets = 9;
   */
  buckets?: StorageBucket;
};

/**
 * Describes the message devkit.v1.FileObject.
 * Use `create(FileObjectSchema)` to create a new message.
 */
export const FileObjectSchema: GenMessage<FileObject> = /*@__PURE__*/
  messageDesc(file_devkit_v1_public_storage, 11);

/**
 * @generated from message devkit.v1.FileListRequest
 */
export type FileListRequest = Message<"devkit.v1.FileListRequest"> & {
  /**
   * @generated from field: string bucket_id = 1;
   */
  bucketId: string;

  /**
   * @generated from field: string query_path = 2;
   */
  queryPath: string;

  /**
   * @generated from field: int32 limit = 3;
   */
  limit: number;

  /**
   * @generated from field: int32 offest = 4;
   */
  offest: number;
};

/**
 * Describes the message devkit.v1.FileListRequest.
 * Use `create(FileListRequestSchema)` to create a new message.
 */
export const FileListRequestSchema: GenMessage<FileListRequest> = /*@__PURE__*/
  messageDesc(file_devkit_v1_public_storage, 12);

/**
 * @generated from message devkit.v1.FileListResponse
 */
export type FileListResponse = Message<"devkit.v1.FileListResponse"> & {
  /**
   * @generated from field: repeated devkit.v1.FileObject files = 1;
   */
  files: FileObject[];
};

/**
 * Describes the message devkit.v1.FileListResponse.
 * Use `create(FileListResponseSchema)` to create a new message.
 */
export const FileListResponseSchema: GenMessage<FileListResponse> = /*@__PURE__*/
  messageDesc(file_devkit_v1_public_storage, 13);

/**
 * @generated from message devkit.v1.FileDeleteRequest
 */
export type FileDeleteRequest = Message<"devkit.v1.FileDeleteRequest"> & {
  /**
   * @generated from field: string bucket_id = 1;
   */
  bucketId: string;

  /**
   * @generated from field: repeated string files_paths = 2;
   */
  filesPaths: string[];
};

/**
 * Describes the message devkit.v1.FileDeleteRequest.
 * Use `create(FileDeleteRequestSchema)` to create a new message.
 */
export const FileDeleteRequestSchema: GenMessage<FileDeleteRequest> = /*@__PURE__*/
  messageDesc(file_devkit_v1_public_storage, 14);

/**
 * @generated from message devkit.v1.FileDeleteResponse
 */
export type FileDeleteResponse = Message<"devkit.v1.FileDeleteResponse"> & {
  /**
   * @generated from field: repeated devkit.v1.FileCreateResponse responses = 1;
   */
  responses: FileCreateResponse[];
};

/**
 * Describes the message devkit.v1.FileDeleteResponse.
 * Use `create(FileDeleteResponseSchema)` to create a new message.
 */
export const FileDeleteResponseSchema: GenMessage<FileDeleteResponse> = /*@__PURE__*/
  messageDesc(file_devkit_v1_public_storage, 15);
