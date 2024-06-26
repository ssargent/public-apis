// @generated by protoc-gen-es v1.9.0 with parameter "target=js+dts"
// @generated from file bbq/collector/v1/bbq.proto (package bbq.collector.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage, Timestamp } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";

/**
 * @generated from message bbq.collector.v1.Sensor
 */
export declare class Sensor extends Message<Sensor> {
  /**
   * @generated from field: string name = 1;
   */
  name: string;

  /**
   * @generated from field: string manufacturer = 2;
   */
  manufacturer: string;

  /**
   * @generated from field: int32 sensor_count = 3;
   */
  sensorCount: number;

  /**
   * @generated from field: string temperature_unit = 4;
   */
  temperatureUnit: string;

  /**
   * @generated from field: string sensor_id = 5;
   */
  sensorId: string;

  constructor(data?: PartialMessage<Sensor>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "bbq.collector.v1.Sensor";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Sensor;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Sensor;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Sensor;

  static equals(a: Sensor | PlainMessage<Sensor> | undefined, b: Sensor | PlainMessage<Sensor> | undefined): boolean;
}

/**
 * @generated from message bbq.collector.v1.SensorReading
 */
export declare class SensorReading extends Message<SensorReading> {
  /**
   * @generated from field: int32 sensor_number = 1;
   */
  sensorNumber: number;

  /**
   * @generated from field: float temperature = 2;
   */
  temperature: number;

  constructor(data?: PartialMessage<SensorReading>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "bbq.collector.v1.SensorReading";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): SensorReading;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): SensorReading;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): SensorReading;

  static equals(a: SensorReading | PlainMessage<SensorReading> | undefined, b: SensorReading | PlainMessage<SensorReading> | undefined): boolean;
}

/**
 * @generated from message bbq.collector.v1.Reading
 */
export declare class Reading extends Message<Reading> {
  /**
   * @generated from field: bbq.collector.v1.Sensor sensor = 1;
   */
  sensor?: Sensor;

  /**
   * @generated from field: repeated bbq.collector.v1.SensorReading readings = 2;
   */
  readings: SensorReading[];

  /**
   * @generated from field: google.protobuf.Timestamp recorded_at = 3;
   */
  recordedAt?: Timestamp;

  constructor(data?: PartialMessage<Reading>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "bbq.collector.v1.Reading";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Reading;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Reading;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Reading;

  static equals(a: Reading | PlainMessage<Reading> | undefined, b: Reading | PlainMessage<Reading> | undefined): boolean;
}

/**
 * @generated from message bbq.collector.v1.Session
 */
export declare class Session extends Message<Session> {
  /**
   * @generated from field: string id = 1;
   */
  id: string;

  /**
   * @generated from field: string name = 2;
   */
  name: string;

  constructor(data?: PartialMessage<Session>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "bbq.collector.v1.Session";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Session;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Session;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Session;

  static equals(a: Session | PlainMessage<Session> | undefined, b: Session | PlainMessage<Session> | undefined): boolean;
}

