// @generated by protoc-gen-es v2.5.1 with parameter "target=js+dts"
// @generated from file bbq/collector/v1/bbq.proto (package bbq.collector.v1, syntax proto3)
/* eslint-disable */

import type { GenFile, GenMessage } from "@bufbuild/protobuf/codegenv2";
import type { Message } from "@bufbuild/protobuf";
import type { Timestamp } from "@bufbuild/protobuf/wkt";

/**
 * Describes the file bbq/collector/v1/bbq.proto.
 */
export declare const file_bbq_collector_v1_bbq: GenFile;

/**
 * @generated from message bbq.collector.v1.Sensor
 */
export declare type Sensor = Message<"bbq.collector.v1.Sensor"> & {
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
};

/**
 * Describes the message bbq.collector.v1.Sensor.
 * Use `create(SensorSchema)` to create a new message.
 */
export declare const SensorSchema: GenMessage<Sensor>;

/**
 * @generated from message bbq.collector.v1.SensorReading
 */
export declare type SensorReading = Message<"bbq.collector.v1.SensorReading"> & {
  /**
   * @generated from field: int32 sensor_number = 1;
   */
  sensorNumber: number;

  /**
   * @generated from field: float temperature = 2;
   */
  temperature: number;
};

/**
 * Describes the message bbq.collector.v1.SensorReading.
 * Use `create(SensorReadingSchema)` to create a new message.
 */
export declare const SensorReadingSchema: GenMessage<SensorReading>;

/**
 * @generated from message bbq.collector.v1.Reading
 */
export declare type Reading = Message<"bbq.collector.v1.Reading"> & {
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
};

/**
 * Describes the message bbq.collector.v1.Reading.
 * Use `create(ReadingSchema)` to create a new message.
 */
export declare const ReadingSchema: GenMessage<Reading>;

/**
 * @generated from message bbq.collector.v1.Session
 */
export declare type Session = Message<"bbq.collector.v1.Session"> & {
  /**
   * @generated from field: string id = 1;
   */
  id: string;

  /**
   * @generated from field: string name = 2;
   */
  name: string;
};

/**
 * Describes the message bbq.collector.v1.Session.
 * Use `create(SessionSchema)` to create a new message.
 */
export declare const SessionSchema: GenMessage<Session>;

