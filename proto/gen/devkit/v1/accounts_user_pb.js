// @generated by protoc-gen-es v2.2.2
// @generated from file devkit/v1/accounts_user.proto (package devkit.v1, syntax proto3)
/* eslint-disable */

import { fileDesc, messageDesc } from "@bufbuild/protobuf/codegenv1";
import { file_buf_validate_validate } from "../../buf/validate/validate_pb";
import { file_devkit_v1_client } from "./client_pb";

/**
 * Describes the file devkit/v1/accounts_user.proto.
 */
export const file_devkit_v1_accounts_user = /*@__PURE__*/
  fileDesc("Ch1kZXZraXQvdjEvYWNjb3VudHNfdXNlci5wcm90bxIJZGV2a2l0LnYxIpICChdVc2VyQ3JlYXRlVXBkYXRlUmVxdWVzdBIPCgd1c2VyX2lkGAEgASgFEh0KCXVzZXJfbmFtZRgCIAEoCUIKukgHcgUQAhjIARIRCgl0ZW5hbnRfaWQYCSABKAUSGwoTdXNlcl9zZWN1cml0eV9sZXZlbBgDIAEoBRIdCgx1c2VyX3R5cGVfaWQYBCABKAVCB7pIBBoCIAASGwoKdXNlcl9waG9uZRgFIAEoCUIHukgEcgIYMhIeCgp1c2VyX2VtYWlsGAYgASgJQgq6SAdyBRjIAWABEh8KDXVzZXJfcGFzc3dvcmQYByABKAlCCLpIBXIDGMgBEhoKBXJvbGVzGAggAygFQgu6SAiSAQUQ9AMYASL5AQoSQWNjb3VudHNTY2hlbWFVc2VyEg8KB3VzZXJfaWQYASABKAUSEQoJdXNlcl9uYW1lGAIgASgJEhEKCXRlbmFudF9pZBgMIAEoBRIbChN1c2VyX3NlY3VyaXR5X2xldmVsGAMgASgFEhQKDHVzZXJfdHlwZV9pZBgEIAEoBRISCgp1c2VyX3Bob25lGAUgASgJEhIKCnVzZXJfZW1haWwYBiABKAkSFQoNdXNlcl9wYXNzd29yZBgHIAEoCRISCgpjcmVhdGVkX2F0GAkgASgJEhIKCnVwZGF0ZWRfYXQYCiABKAkSEgoKZGVsZXRlZF9hdBgLIAEoCSJHChhVc2VyQ3JlYXRlVXBkYXRlUmVzcG9uc2USKwoEdXNlchgBIAEoCzIdLmRldmtpdC52MS5BY2NvdW50c1NjaGVtYVVzZXIiKwoYVXNlckRlbGV0ZVJlc3RvcmVSZXF1ZXN0Eg8KB3JlY29yZHMYASADKAUiSwoZVXNlckRlbGV0ZVJlc3RvcmVSZXNwb25zZRIuCgdyZWNvcmRzGAEgAygLMh0uZGV2a2l0LnYxLkFjY291bnRzU2NoZW1hVXNlciImChFVc2VyRGVsZXRlUmVxdWVzdBIRCglyZWNvcmRfaWQYASABKAUiQwoSVXNlckRlbGV0ZVJlc3BvbnNlEi0KBnJlY29yZBgBIAEoCzIdLmRldmtpdC52MS5BY2NvdW50c1NjaGVtYVVzZXIiEQoPVXNlckxpc3RSZXF1ZXN0IqgBChBVc2VyTGlzdFJlc3BvbnNlEi4KB3JlY29yZHMYASADKAsyHS5kZXZraXQudjEuQWNjb3VudHNTY2hlbWFVc2VyEjYKD2RlbGV0ZWRfcmVjb3JkcxgCIAMoCzIdLmRldmtpdC52MS5BY2NvdW50c1NjaGVtYVVzZXISLAoHb3B0aW9ucxgDIAEoCzIbLmRldmtpdC52MS5BdmFpbGFibGVPcHRpb25zIhYKFFVzZXJMaXN0SW5wdXRSZXF1ZXN0IkYKFVVzZXJMaXN0SW5wdXRSZXNwb25zZRItCgdvcHRpb25zGAEgAygLMhwuZGV2a2l0LnYxLlNlbGVjdElucHV0T3B0aW9uIi0KGFVzZXJGaW5kRm9yVXBkYXRlUmVxdWVzdBIRCglyZWNvcmRfaWQYASABKAUiUAoZVXNlckZpbmRGb3JVcGRhdGVSZXNwb25zZRIzCgdyZXF1ZXN0GAEgASgLMiIuZGV2a2l0LnYxLlVzZXJDcmVhdGVVcGRhdGVSZXF1ZXN0QqYBCg1jb20uZGV2a2l0LnYxQhFBY2NvdW50c1VzZXJQcm90b1ABWj1naXRodWIuY29tL2Rhcndpc2hkZXYvZGV2a2l0LWFwaS9wcm90b19nZW4vZGV2a2l0L3YxO2RldmtpdHYxogIDRFhYqgIJRGV2a2l0LlYxygIJRGV2a2l0XFYx4gIVRGV2a2l0XFYxXEdQQk1ldGFkYXRh6gIKRGV2a2l0OjpWMWIGcHJvdG8z", [file_buf_validate_validate, file_devkit_v1_client]);

/**
 * Describes the message devkit.v1.UserCreateUpdateRequest.
 * Use `create(UserCreateUpdateRequestSchema)` to create a new message.
 */
export const UserCreateUpdateRequestSchema = /*@__PURE__*/
  messageDesc(file_devkit_v1_accounts_user, 0);

/**
 * Describes the message devkit.v1.AccountsSchemaUser.
 * Use `create(AccountsSchemaUserSchema)` to create a new message.
 */
export const AccountsSchemaUserSchema = /*@__PURE__*/
  messageDesc(file_devkit_v1_accounts_user, 1);

/**
 * Describes the message devkit.v1.UserCreateUpdateResponse.
 * Use `create(UserCreateUpdateResponseSchema)` to create a new message.
 */
export const UserCreateUpdateResponseSchema = /*@__PURE__*/
  messageDesc(file_devkit_v1_accounts_user, 2);

/**
 * Describes the message devkit.v1.UserDeleteRestoreRequest.
 * Use `create(UserDeleteRestoreRequestSchema)` to create a new message.
 */
export const UserDeleteRestoreRequestSchema = /*@__PURE__*/
  messageDesc(file_devkit_v1_accounts_user, 3);

/**
 * Describes the message devkit.v1.UserDeleteRestoreResponse.
 * Use `create(UserDeleteRestoreResponseSchema)` to create a new message.
 */
export const UserDeleteRestoreResponseSchema = /*@__PURE__*/
  messageDesc(file_devkit_v1_accounts_user, 4);

/**
 * Describes the message devkit.v1.UserDeleteRequest.
 * Use `create(UserDeleteRequestSchema)` to create a new message.
 */
export const UserDeleteRequestSchema = /*@__PURE__*/
  messageDesc(file_devkit_v1_accounts_user, 5);

/**
 * Describes the message devkit.v1.UserDeleteResponse.
 * Use `create(UserDeleteResponseSchema)` to create a new message.
 */
export const UserDeleteResponseSchema = /*@__PURE__*/
  messageDesc(file_devkit_v1_accounts_user, 6);

/**
 * Describes the message devkit.v1.UserListRequest.
 * Use `create(UserListRequestSchema)` to create a new message.
 */
export const UserListRequestSchema = /*@__PURE__*/
  messageDesc(file_devkit_v1_accounts_user, 7);

/**
 * Describes the message devkit.v1.UserListResponse.
 * Use `create(UserListResponseSchema)` to create a new message.
 */
export const UserListResponseSchema = /*@__PURE__*/
  messageDesc(file_devkit_v1_accounts_user, 8);

/**
 * Describes the message devkit.v1.UserListInputRequest.
 * Use `create(UserListInputRequestSchema)` to create a new message.
 */
export const UserListInputRequestSchema = /*@__PURE__*/
  messageDesc(file_devkit_v1_accounts_user, 9);

/**
 * Describes the message devkit.v1.UserListInputResponse.
 * Use `create(UserListInputResponseSchema)` to create a new message.
 */
export const UserListInputResponseSchema = /*@__PURE__*/
  messageDesc(file_devkit_v1_accounts_user, 10);

/**
 * Describes the message devkit.v1.UserFindForUpdateRequest.
 * Use `create(UserFindForUpdateRequestSchema)` to create a new message.
 */
export const UserFindForUpdateRequestSchema = /*@__PURE__*/
  messageDesc(file_devkit_v1_accounts_user, 11);

/**
 * Describes the message devkit.v1.UserFindForUpdateResponse.
 * Use `create(UserFindForUpdateResponseSchema)` to create a new message.
 */
export const UserFindForUpdateResponseSchema = /*@__PURE__*/
  messageDesc(file_devkit_v1_accounts_user, 12);
