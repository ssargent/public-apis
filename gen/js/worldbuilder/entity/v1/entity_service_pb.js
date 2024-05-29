// @generated by protoc-gen-es v1.9.0 with parameter "target=js+dts"
// @generated from file worldbuilder/entity/v1/entity_service.proto (package worldbuilder.entity.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { proto3 } from "@bufbuild/protobuf";
import { Entity, EntityInput, Filter, Type } from "./entity_pb.js";

/**
 * @generated from message worldbuilder.entity.v1.GetEntityRequest
 */
export const GetEntityRequest = /*@__PURE__*/ proto3.makeMessageType(
  "worldbuilder.entity.v1.GetEntityRequest",
  () => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message worldbuilder.entity.v1.GetEntityResponse
 */
export const GetEntityResponse = /*@__PURE__*/ proto3.makeMessageType(
  "worldbuilder.entity.v1.GetEntityResponse",
  () => [
    { no: 1, name: "entity", kind: "message", T: Entity },
  ],
);

/**
 * @generated from message worldbuilder.entity.v1.GetEntitiesRequest
 */
export const GetEntitiesRequest = /*@__PURE__*/ proto3.makeMessageType(
  "worldbuilder.entity.v1.GetEntitiesRequest",
  () => [
    { no: 1, name: "criteria", kind: "message", T: Filter },
  ],
);

/**
 * @generated from message worldbuilder.entity.v1.GetEntitiesResponse
 */
export const GetEntitiesResponse = /*@__PURE__*/ proto3.makeMessageType(
  "worldbuilder.entity.v1.GetEntitiesResponse",
  () => [
    { no: 1, name: "criteria", kind: "message", T: Filter },
    { no: 2, name: "entities", kind: "message", T: Entity, repeated: true },
  ],
);

/**
 * @generated from message worldbuilder.entity.v1.CreateTypeRequest
 */
export const CreateTypeRequest = /*@__PURE__*/ proto3.makeMessageType(
  "worldbuilder.entity.v1.CreateTypeRequest",
  () => [
    { no: 1, name: "type", kind: "message", T: Type },
  ],
);

/**
 * @generated from message worldbuilder.entity.v1.CreateTypeResponse
 */
export const CreateTypeResponse = /*@__PURE__*/ proto3.makeMessageType(
  "worldbuilder.entity.v1.CreateTypeResponse",
  () => [
    { no: 1, name: "type", kind: "message", T: Type },
  ],
);

/**
 * @generated from message worldbuilder.entity.v1.CreateEntityRequest
 */
export const CreateEntityRequest = /*@__PURE__*/ proto3.makeMessageType(
  "worldbuilder.entity.v1.CreateEntityRequest",
  () => [
    { no: 1, name: "entity", kind: "message", T: EntityInput },
  ],
);

/**
 * @generated from message worldbuilder.entity.v1.CreateEntityResponse
 */
export const CreateEntityResponse = /*@__PURE__*/ proto3.makeMessageType(
  "worldbuilder.entity.v1.CreateEntityResponse",
  () => [
    { no: 1, name: "entity", kind: "message", T: Entity },
  ],
);
