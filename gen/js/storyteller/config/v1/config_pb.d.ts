// @generated by protoc-gen-es v1.9.0 with parameter "target=js+dts"
// @generated from file storyteller/config/v1/config.proto (package storyteller.config.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";

/**
 * @generated from message storyteller.config.v1.ConfigEntry
 */
export declare class ConfigEntry extends Message<ConfigEntry> {
  /**
   * @generated from field: string id = 1;
   */
  id: string;

  /**
   * @generated from field: string key = 2;
   */
  key: string;

  /**
   * @generated from field: bytes value = 3;
   */
  value: Uint8Array;

  /**
   * @generated from field: string data_type = 4;
   */
  dataType: string;

  /**
   * @generated from field: int64 created_at = 5;
   */
  createdAt: bigint;

  /**
   * @generated from field: int64 updated_at = 6;
   */
  updatedAt: bigint;

  constructor(data?: PartialMessage<ConfigEntry>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "storyteller.config.v1.ConfigEntry";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ConfigEntry;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ConfigEntry;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ConfigEntry;

  static equals(a: ConfigEntry | PlainMessage<ConfigEntry> | undefined, b: ConfigEntry | PlainMessage<ConfigEntry> | undefined): boolean;
}

