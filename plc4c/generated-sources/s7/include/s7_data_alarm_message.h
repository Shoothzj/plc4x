/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

#ifndef PLC4C_S7_READ_WRITE_S7_DATA_ALARM_MESSAGE_H_
#define PLC4C_S7_READ_WRITE_S7_DATA_ALARM_MESSAGE_H_

#include <stdbool.h>
#include <stdint.h>
#include <plc4c/driver_s7_static.h>
#include <plc4c/spi/read_buffer.h>
#include <plc4c/spi/write_buffer.h>
#include <plc4c/utils/list.h>
#include "data_transport_error_code.h"
#include "data_transport_size.h"
#include "syntax_id_type.h"
#include "alarm_type.h"
#include "s7_data_alarm_message.h"
#include "query_type.h"

// Code generated by code-generation. DO NOT EDIT.

#ifdef __cplusplus
extern "C" {
#endif


// Structure used to contain the discriminator values for discriminated types using this as a parent
struct plc4c_s7_read_write_s7_data_alarm_message_discriminator {
  uint8_t cpuFunctionType;
};
typedef struct plc4c_s7_read_write_s7_data_alarm_message_discriminator plc4c_s7_read_write_s7_data_alarm_message_discriminator;

// Enum assigning each sub-type an individual id.
enum plc4c_s7_read_write_s7_data_alarm_message_type {
  plc4c_s7_read_write_s7_data_alarm_message_type_plc4c_s7_read_write_s7_message_object_request = 0,
  plc4c_s7_read_write_s7_data_alarm_message_type_plc4c_s7_read_write_s7_message_object_response = 1};
typedef enum plc4c_s7_read_write_s7_data_alarm_message_type plc4c_s7_read_write_s7_data_alarm_message_type;

// Function to get the discriminator values for a given type.
plc4c_s7_read_write_s7_data_alarm_message_discriminator plc4c_s7_read_write_s7_data_alarm_message_get_discriminator(plc4c_s7_read_write_s7_data_alarm_message_type type);

// Constant values.
uint8_t PLC4C_S7_READ_WRITE_S7_DATA_ALARM_MESSAGE_FUNCTION_ID();
uint8_t PLC4C_S7_READ_WRITE_S7_MESSAGE_OBJECT_REQUEST_LENGTH();
uint8_t PLC4C_S7_READ_WRITE_S7_DATA_ALARM_MESSAGE_NUMBER_MESSAGE_OBJ();
uint8_t PLC4C_S7_READ_WRITE_S7_MESSAGE_OBJECT_REQUEST_VARIABLE_SPEC();

struct plc4c_s7_read_write_s7_data_alarm_message {
  /* This is an abstract type so this property saves the type of this typed union */
  plc4c_s7_read_write_s7_data_alarm_message_type _type;
  /* Properties */
  uint8_t function_id;
  uint8_t number_message_obj;
  union {
    struct { /* S7MessageObjectRequest */
      plc4c_s7_read_write_syntax_id_type s7_message_object_request_syntax_id;
      plc4c_s7_read_write_query_type s7_message_object_request_query_type;
      plc4c_s7_read_write_alarm_type s7_message_object_request_alarm_type;
    };
    struct { /* S7MessageObjectResponse */
      plc4c_s7_read_write_data_transport_error_code s7_message_object_response_return_code;
      plc4c_s7_read_write_data_transport_size s7_message_object_response_transport_size;
    };
  };
};
typedef struct plc4c_s7_read_write_s7_data_alarm_message plc4c_s7_read_write_s7_data_alarm_message;

// Create an empty NULL-struct
plc4c_s7_read_write_s7_data_alarm_message plc4c_s7_read_write_s7_data_alarm_message_null();

plc4c_return_code plc4c_s7_read_write_s7_data_alarm_message_parse(plc4c_spi_read_buffer* readBuffer, uint8_t cpuFunctionType, plc4c_s7_read_write_s7_data_alarm_message** message);

plc4c_return_code plc4c_s7_read_write_s7_data_alarm_message_serialize(plc4c_spi_write_buffer* writeBuffer, plc4c_s7_read_write_s7_data_alarm_message* message);

uint16_t plc4c_s7_read_write_s7_data_alarm_message_length_in_bytes(plc4c_s7_read_write_s7_data_alarm_message* message);

uint16_t plc4c_s7_read_write_s7_data_alarm_message_length_in_bits(plc4c_s7_read_write_s7_data_alarm_message* message);

#ifdef __cplusplus
}
#endif
#endif  // PLC4C_S7_READ_WRITE_S7_DATA_ALARM_MESSAGE_H_
