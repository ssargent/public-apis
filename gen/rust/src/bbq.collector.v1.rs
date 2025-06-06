// @generated
// This file is @generated by prost-build.
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct Sensor {
    #[prost(string, tag="1")]
    pub name: ::prost::alloc::string::String,
    #[prost(string, tag="2")]
    pub manufacturer: ::prost::alloc::string::String,
    #[prost(int32, tag="3")]
    pub sensor_count: i32,
    #[prost(string, tag="4")]
    pub temperature_unit: ::prost::alloc::string::String,
    #[prost(string, tag="5")]
    pub sensor_id: ::prost::alloc::string::String,
}
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, Copy, PartialEq, ::prost::Message)]
pub struct SensorReading {
    #[prost(int32, tag="1")]
    pub sensor_number: i32,
    #[prost(float, tag="2")]
    pub temperature: f32,
}
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct Reading {
    #[prost(message, optional, tag="1")]
    pub sensor: ::core::option::Option<Sensor>,
    #[prost(message, repeated, tag="2")]
    pub readings: ::prost::alloc::vec::Vec<SensorReading>,
    #[prost(message, optional, tag="3")]
    pub recorded_at: ::core::option::Option<::pbjson_types::Timestamp>,
}
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct Session {
    #[prost(string, tag="1")]
    pub id: ::prost::alloc::string::String,
    #[prost(string, tag="2")]
    pub name: ::prost::alloc::string::String,
}
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct RecordRequest {
    #[prost(message, optional, tag="1")]
    pub reading: ::core::option::Option<Reading>,
}
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct RecordResponse {
    #[prost(string, tag="1")]
    pub session_id: ::prost::alloc::string::String,
    #[prost(message, optional, tag="2")]
    pub recorded_at: ::core::option::Option<::pbjson_types::Timestamp>,
}
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct SessionRequest {
    #[prost(string, tag="1")]
    pub name: ::prost::alloc::string::String,
}
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct SessionResponse {
    #[prost(message, optional, tag="1")]
    pub session: ::core::option::Option<Session>,
}
include!("bbq.collector.v1.tonic.rs");
// @@protoc_insertion_point(module)