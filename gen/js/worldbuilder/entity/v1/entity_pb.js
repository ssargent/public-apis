// @generated by protoc-gen-es v1.9.0 with parameter "target=js+dts"
// @generated from file worldbuilder/entity/v1/entity.proto (package worldbuilder.entity.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { proto3 } from "@bufbuild/protobuf";

/**
 * @generated from message worldbuilder.entity.v1.AttributeDefinition
 */
export const AttributeDefinition = /*@__PURE__*/ proto3.makeMessageType(
  "worldbuilder.entity.v1.AttributeDefinition",
  () => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "wbatn", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "attribute_name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "label", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 5, name: "data_type", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 6, name: "created_at", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
    { no: 7, name: "updated_at", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
  ],
);

/**
 * @generated from message worldbuilder.entity.v1.TypeAttribute
 */
export const TypeAttribute = /*@__PURE__*/ proto3.makeMessageType(
  "worldbuilder.entity.v1.TypeAttribute",
  () => [
    { no: 1, name: "wbatn", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "attribute_name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "ordinal", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 4, name: "is_required", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
  ],
);

/**
 * @generated from message worldbuilder.entity.v1.TypeParent
 */
export const TypeParent = /*@__PURE__*/ proto3.makeMessageType(
  "worldbuilder.entity.v1.TypeParent",
  () => [
    { no: 1, name: "type_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "wbtn", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "type_name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "type_description", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message worldbuilder.entity.v1.Type
 */
export const Type = /*@__PURE__*/ proto3.makeMessageType(
  "worldbuilder.entity.v1.Type",
  () => [
    { no: 1, name: "type_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "wbtn", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "type_name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "type_description", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 5, name: "created_at", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
    { no: 6, name: "updated_at", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
    { no: 7, name: "parent", kind: "message", T: TypeParent },
    { no: 8, name: "attributes", kind: "message", T: TypeAttribute, repeated: true },
  ],
);

/**
 * @generated from message worldbuilder.entity.v1.EntityAttribute
 */
export const EntityAttribute = /*@__PURE__*/ proto3.makeMessageType(
  "worldbuilder.entity.v1.EntityAttribute",
  () => [
    { no: 1, name: "name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "type", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "value", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message worldbuilder.entity.v1.EntityParent
 */
export const EntityParent = /*@__PURE__*/ proto3.makeMessageType(
  "worldbuilder.entity.v1.EntityParent",
  () => [
    { no: 1, name: "entity_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "entity_name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "resource_name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "type_name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message worldbuilder.entity.v1.EntityType
 */
export const EntityType = /*@__PURE__*/ proto3.makeMessageType(
  "worldbuilder.entity.v1.EntityType",
  () => [
    { no: 1, name: "type_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "type_name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message worldbuilder.entity.v1.Entity
 */
export const Entity = /*@__PURE__*/ proto3.makeMessageType(
  "worldbuilder.entity.v1.Entity",
  () => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "resource_name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "description", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 5, name: "created_at", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 6, name: "updated_at", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 7, name: "parent", kind: "message", T: EntityParent },
    { no: 8, name: "type", kind: "message", T: EntityType },
    { no: 9, name: "attributes", kind: "message", T: EntityAttribute, repeated: true },
  ],
);

/**
 * @generated from message worldbuilder.entity.v1.EntityInput
 */
export const EntityInput = /*@__PURE__*/ proto3.makeMessageType(
  "worldbuilder.entity.v1.EntityInput",
  () => [
    { no: 1, name: "resource_name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "description", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "parent_wbrn", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 5, name: "type_name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 6, name: "attributes", kind: "message", T: EntityAttribute, repeated: true },
  ],
);

/**
 * @generated from message worldbuilder.entity.v1.Filter
 */
export const Filter = /*@__PURE__*/ proto3.makeMessageType(
  "worldbuilder.entity.v1.Filter",
  () => [
    { no: 1, name: "type_name", kind: "scalar", T: 9 /* ScalarType.STRING */, opt: true },
    { no: 2, name: "parent_wbrn", kind: "scalar", T: 9 /* ScalarType.STRING */, opt: true },
  ],
);

