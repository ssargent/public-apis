// @generated by protoc-gen-es v1.9.0 with parameter "target=js+dts"
// @generated from file alexandria/catalogue/v1/catalogue_service.proto (package alexandria.catalogue.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";
import type { Item, ItemInput, Room, RoomInput, Shelf, ShelfInput } from "./catalogue_pb.js";

/**
 * @generated from message alexandria.catalogue.v1.ErrorResponse
 */
export declare class ErrorResponse extends Message<ErrorResponse> {
  /**
   * @generated from field: string code = 1;
   */
  code: string;

  /**
   * @generated from field: string message = 2;
   */
  message: string;

  constructor(data?: PartialMessage<ErrorResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "alexandria.catalogue.v1.ErrorResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ErrorResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ErrorResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ErrorResponse;

  static equals(a: ErrorResponse | PlainMessage<ErrorResponse> | undefined, b: ErrorResponse | PlainMessage<ErrorResponse> | undefined): boolean;
}

/**
 * @generated from message alexandria.catalogue.v1.CreateRoomRequest
 */
export declare class CreateRoomRequest extends Message<CreateRoomRequest> {
  /**
   * @generated from field: repeated alexandria.catalogue.v1.RoomInput rooms = 1;
   */
  rooms: RoomInput[];

  constructor(data?: PartialMessage<CreateRoomRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "alexandria.catalogue.v1.CreateRoomRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CreateRoomRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CreateRoomRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CreateRoomRequest;

  static equals(a: CreateRoomRequest | PlainMessage<CreateRoomRequest> | undefined, b: CreateRoomRequest | PlainMessage<CreateRoomRequest> | undefined): boolean;
}

/**
 * @generated from message alexandria.catalogue.v1.Rooms
 */
export declare class Rooms extends Message<Rooms> {
  /**
   * @generated from field: repeated alexandria.catalogue.v1.Room rooms = 1;
   */
  rooms: Room[];

  constructor(data?: PartialMessage<Rooms>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "alexandria.catalogue.v1.Rooms";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Rooms;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Rooms;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Rooms;

  static equals(a: Rooms | PlainMessage<Rooms> | undefined, b: Rooms | PlainMessage<Rooms> | undefined): boolean;
}

/**
 * @generated from message alexandria.catalogue.v1.CreateRoomResponse
 */
export declare class CreateRoomResponse extends Message<CreateRoomResponse> {
  /**
   * @generated from oneof alexandria.catalogue.v1.CreateRoomResponse.result
   */
  result: {
    /**
     * @generated from field: alexandria.catalogue.v1.Rooms rooms = 1;
     */
    value: Rooms;
    case: "rooms";
  } | {
    /**
     * @generated from field: alexandria.catalogue.v1.ErrorResponse error = 2;
     */
    value: ErrorResponse;
    case: "error";
  } | { case: undefined; value?: undefined };

  constructor(data?: PartialMessage<CreateRoomResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "alexandria.catalogue.v1.CreateRoomResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CreateRoomResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CreateRoomResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CreateRoomResponse;

  static equals(a: CreateRoomResponse | PlainMessage<CreateRoomResponse> | undefined, b: CreateRoomResponse | PlainMessage<CreateRoomResponse> | undefined): boolean;
}

/**
 * @generated from message alexandria.catalogue.v1.CreateItemRequest
 */
export declare class CreateItemRequest extends Message<CreateItemRequest> {
  /**
   * @generated from field: repeated alexandria.catalogue.v1.ItemInput items = 1;
   */
  items: ItemInput[];

  constructor(data?: PartialMessage<CreateItemRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "alexandria.catalogue.v1.CreateItemRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CreateItemRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CreateItemRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CreateItemRequest;

  static equals(a: CreateItemRequest | PlainMessage<CreateItemRequest> | undefined, b: CreateItemRequest | PlainMessage<CreateItemRequest> | undefined): boolean;
}

/**
 * @generated from message alexandria.catalogue.v1.Items
 */
export declare class Items extends Message<Items> {
  /**
   * @generated from field: repeated alexandria.catalogue.v1.Item items = 1;
   */
  items: Item[];

  constructor(data?: PartialMessage<Items>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "alexandria.catalogue.v1.Items";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Items;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Items;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Items;

  static equals(a: Items | PlainMessage<Items> | undefined, b: Items | PlainMessage<Items> | undefined): boolean;
}

/**
 * @generated from message alexandria.catalogue.v1.CreateItemResponse
 */
export declare class CreateItemResponse extends Message<CreateItemResponse> {
  /**
   * @generated from oneof alexandria.catalogue.v1.CreateItemResponse.result
   */
  result: {
    /**
     * @generated from field: alexandria.catalogue.v1.Items items = 1;
     */
    value: Items;
    case: "items";
  } | {
    /**
     * @generated from field: alexandria.catalogue.v1.ErrorResponse error = 2;
     */
    value: ErrorResponse;
    case: "error";
  } | { case: undefined; value?: undefined };

  constructor(data?: PartialMessage<CreateItemResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "alexandria.catalogue.v1.CreateItemResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CreateItemResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CreateItemResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CreateItemResponse;

  static equals(a: CreateItemResponse | PlainMessage<CreateItemResponse> | undefined, b: CreateItemResponse | PlainMessage<CreateItemResponse> | undefined): boolean;
}

/**
 * @generated from message alexandria.catalogue.v1.CreateShelfRequest
 */
export declare class CreateShelfRequest extends Message<CreateShelfRequest> {
  /**
   * @generated from field: repeated alexandria.catalogue.v1.ShelfInput shelves = 1;
   */
  shelves: ShelfInput[];

  constructor(data?: PartialMessage<CreateShelfRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "alexandria.catalogue.v1.CreateShelfRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CreateShelfRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CreateShelfRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CreateShelfRequest;

  static equals(a: CreateShelfRequest | PlainMessage<CreateShelfRequest> | undefined, b: CreateShelfRequest | PlainMessage<CreateShelfRequest> | undefined): boolean;
}

/**
 * @generated from message alexandria.catalogue.v1.Shelves
 */
export declare class Shelves extends Message<Shelves> {
  /**
   * @generated from field: repeated alexandria.catalogue.v1.Shelf shelves = 1;
   */
  shelves: Shelf[];

  constructor(data?: PartialMessage<Shelves>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "alexandria.catalogue.v1.Shelves";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Shelves;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Shelves;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Shelves;

  static equals(a: Shelves | PlainMessage<Shelves> | undefined, b: Shelves | PlainMessage<Shelves> | undefined): boolean;
}

/**
 * @generated from message alexandria.catalogue.v1.CreateShelfResponse
 */
export declare class CreateShelfResponse extends Message<CreateShelfResponse> {
  /**
   * @generated from oneof alexandria.catalogue.v1.CreateShelfResponse.result
   */
  result: {
    /**
     * @generated from field: alexandria.catalogue.v1.Shelves shelves = 1;
     */
    value: Shelves;
    case: "shelves";
  } | {
    /**
     * @generated from field: alexandria.catalogue.v1.ErrorResponse error = 2;
     */
    value: ErrorResponse;
    case: "error";
  } | { case: undefined; value?: undefined };

  constructor(data?: PartialMessage<CreateShelfResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "alexandria.catalogue.v1.CreateShelfResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CreateShelfResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CreateShelfResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CreateShelfResponse;

  static equals(a: CreateShelfResponse | PlainMessage<CreateShelfResponse> | undefined, b: CreateShelfResponse | PlainMessage<CreateShelfResponse> | undefined): boolean;
}

/**
 * @generated from message alexandria.catalogue.v1.MoveItemRequest
 */
export declare class MoveItemRequest extends Message<MoveItemRequest> {
  /**
   * @generated from field: string item_id = 1;
   */
  itemId: string;

  /**
   * @generated from field: string shelf_id = 2;
   */
  shelfId: string;

  constructor(data?: PartialMessage<MoveItemRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "alexandria.catalogue.v1.MoveItemRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): MoveItemRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): MoveItemRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): MoveItemRequest;

  static equals(a: MoveItemRequest | PlainMessage<MoveItemRequest> | undefined, b: MoveItemRequest | PlainMessage<MoveItemRequest> | undefined): boolean;
}

/**
 * @generated from message alexandria.catalogue.v1.MoveItemResponse
 */
export declare class MoveItemResponse extends Message<MoveItemResponse> {
  /**
   * @generated from oneof alexandria.catalogue.v1.MoveItemResponse.result
   */
  result: {
    /**
     * @generated from field: alexandria.catalogue.v1.Item item = 1;
     */
    value: Item;
    case: "item";
  } | {
    /**
     * @generated from field: alexandria.catalogue.v1.ErrorResponse error = 2;
     */
    value: ErrorResponse;
    case: "error";
  } | { case: undefined; value?: undefined };

  constructor(data?: PartialMessage<MoveItemResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "alexandria.catalogue.v1.MoveItemResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): MoveItemResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): MoveItemResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): MoveItemResponse;

  static equals(a: MoveItemResponse | PlainMessage<MoveItemResponse> | undefined, b: MoveItemResponse | PlainMessage<MoveItemResponse> | undefined): boolean;
}

/**
 * @generated from message alexandria.catalogue.v1.Search
 */
export declare class Search extends Message<Search> {
  /**
   * @generated from field: string query = 1;
   */
  query: string;

  /**
   * @generated from field: string author = 2;
   */
  author: string;

  /**
   * @generated from field: repeated string categories = 3;
   */
  categories: string[];

  constructor(data?: PartialMessage<Search>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "alexandria.catalogue.v1.Search";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Search;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Search;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Search;

  static equals(a: Search | PlainMessage<Search> | undefined, b: Search | PlainMessage<Search> | undefined): boolean;
}

/**
 * @generated from message alexandria.catalogue.v1.SearchRequest
 */
export declare class SearchRequest extends Message<SearchRequest> {
  /**
   * @generated from field: alexandria.catalogue.v1.Search search = 1;
   */
  search?: Search;

  constructor(data?: PartialMessage<SearchRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "alexandria.catalogue.v1.SearchRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): SearchRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): SearchRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): SearchRequest;

  static equals(a: SearchRequest | PlainMessage<SearchRequest> | undefined, b: SearchRequest | PlainMessage<SearchRequest> | undefined): boolean;
}

/**
 * @generated from message alexandria.catalogue.v1.SearchResponse
 */
export declare class SearchResponse extends Message<SearchResponse> {
  /**
   * @generated from oneof alexandria.catalogue.v1.SearchResponse.result
   */
  result: {
    /**
     * @generated from field: alexandria.catalogue.v1.Items items = 1;
     */
    value: Items;
    case: "items";
  } | {
    /**
     * @generated from field: alexandria.catalogue.v1.ErrorResponse error = 2;
     */
    value: ErrorResponse;
    case: "error";
  } | { case: undefined; value?: undefined };

  constructor(data?: PartialMessage<SearchResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "alexandria.catalogue.v1.SearchResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): SearchResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): SearchResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): SearchResponse;

  static equals(a: SearchResponse | PlainMessage<SearchResponse> | undefined, b: SearchResponse | PlainMessage<SearchResponse> | undefined): boolean;
}

