// @generated by protoc-gen-connect-es v1.4.0 with parameter "target=js+dts"
// @generated from file shopping/list/v1/service.proto (package shopping.list.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { AddToListRequest, AddToListResponse, CloneListRequest, CloneListResponse, CreateListRequest, CreateListResponse, GetListRequest, GetListResponse, MergeListRequest, MergeListResponse, RemoveFromListRequest, RemoveFromListResponse } from "./service_pb.js";
import { MethodKind } from "@bufbuild/protobuf";

/**
 * @generated from service shopping.list.v1.ShoppingListService
 */
export declare const ShoppingListService: {
  readonly typeName: "shopping.list.v1.ShoppingListService",
  readonly methods: {
    /**
     * @generated from rpc shopping.list.v1.ShoppingListService.CreateList
     */
    readonly createList: {
      readonly name: "CreateList",
      readonly I: typeof CreateListRequest,
      readonly O: typeof CreateListResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc shopping.list.v1.ShoppingListService.CloneList
     */
    readonly cloneList: {
      readonly name: "CloneList",
      readonly I: typeof CloneListRequest,
      readonly O: typeof CloneListResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc shopping.list.v1.ShoppingListService.AddToList
     */
    readonly addToList: {
      readonly name: "AddToList",
      readonly I: typeof AddToListRequest,
      readonly O: typeof AddToListResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc shopping.list.v1.ShoppingListService.RemoveFromList
     */
    readonly removeFromList: {
      readonly name: "RemoveFromList",
      readonly I: typeof RemoveFromListRequest,
      readonly O: typeof RemoveFromListResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc shopping.list.v1.ShoppingListService.GetList
     */
    readonly getList: {
      readonly name: "GetList",
      readonly I: typeof GetListRequest,
      readonly O: typeof GetListResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc shopping.list.v1.ShoppingListService.MergeList
     */
    readonly mergeList: {
      readonly name: "MergeList",
      readonly I: typeof MergeListRequest,
      readonly O: typeof MergeListResponse,
      readonly kind: MethodKind.Unary,
    },
  }
};

