// @generated by protoc-gen-es v1.9.0 with parameter "target=js+dts"
// @generated from file bbq/collector/v1/bbq.proto (package bbq.collector.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { proto3, Timestamp } from "@bufbuild/protobuf";

/**
 * @generated from message bbq.collector.v1.Sensor
 */
export const Sensor = /*@__PURE__*/ proto3.makeMessageType(
  "bbq.collector.v1.Sensor",
  () => [
    { no: 1, name: "name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "manufacturer", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "sensor_count", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 4, name: "temperature_unit", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 5, name: "sensor_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message bbq.collector.v1.SensorReading
 */
export const SensorReading = /*@__PURE__*/ proto3.makeMessageType(
  "bbq.collector.v1.SensorReading",
  () => [
    { no: 1, name: "sensor_number", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 2, name: "temperature", kind: "scalar", T: 2 /* ScalarType.FLOAT */ },
  ],
);

/**
 * @generated from message bbq.collector.v1.Reading
 */
export const Reading = /*@__PURE__*/ proto3.makeMessageType(
  "bbq.collector.v1.Reading",
  () => [
    { no: 1, name: "sensor", kind: "message", T: Sensor },
    { no: 2, name: "readings", kind: "message", T: SensorReading, repeated: true },
    { no: 3, name: "recorded_at", kind: "message", T: Timestamp },
  ],
);

/**
 * @generated from message bbq.collector.v1.Session
 */
export const Session = /*@__PURE__*/ proto3.makeMessageType(
  "bbq.collector.v1.Session",
  () => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

