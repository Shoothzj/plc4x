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

package model

import (
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

type ModeTransitionType uint8

type IModeTransitionType interface {
	Serialize(writeBuffer utils.WriteBuffer) error
}

const (
	ModeTransitionType_STOP         ModeTransitionType = 0x00
	ModeTransitionType_WARM_RESTART ModeTransitionType = 0x01
	ModeTransitionType_RUN          ModeTransitionType = 0x02
	ModeTransitionType_HOT_RESTART  ModeTransitionType = 0x03
	ModeTransitionType_HOLD         ModeTransitionType = 0x04
	ModeTransitionType_COLD_RESTART ModeTransitionType = 0x06
	ModeTransitionType_RUN_R        ModeTransitionType = 0x09
	ModeTransitionType_LINK_UP      ModeTransitionType = 0x11
	ModeTransitionType_UPDATE       ModeTransitionType = 0x12
)

var ModeTransitionTypeValues []ModeTransitionType

func init() {
	_ = errors.New
	ModeTransitionTypeValues = []ModeTransitionType{
		ModeTransitionType_STOP,
		ModeTransitionType_WARM_RESTART,
		ModeTransitionType_RUN,
		ModeTransitionType_HOT_RESTART,
		ModeTransitionType_HOLD,
		ModeTransitionType_COLD_RESTART,
		ModeTransitionType_RUN_R,
		ModeTransitionType_LINK_UP,
		ModeTransitionType_UPDATE,
	}
}

func ModeTransitionTypeByValue(value uint8) ModeTransitionType {
	switch value {
	case 0x00:
		return ModeTransitionType_STOP
	case 0x01:
		return ModeTransitionType_WARM_RESTART
	case 0x02:
		return ModeTransitionType_RUN
	case 0x03:
		return ModeTransitionType_HOT_RESTART
	case 0x04:
		return ModeTransitionType_HOLD
	case 0x06:
		return ModeTransitionType_COLD_RESTART
	case 0x09:
		return ModeTransitionType_RUN_R
	case 0x11:
		return ModeTransitionType_LINK_UP
	case 0x12:
		return ModeTransitionType_UPDATE
	}
	return 0
}

func ModeTransitionTypeByName(value string) ModeTransitionType {
	switch value {
	case "STOP":
		return ModeTransitionType_STOP
	case "WARM_RESTART":
		return ModeTransitionType_WARM_RESTART
	case "RUN":
		return ModeTransitionType_RUN
	case "HOT_RESTART":
		return ModeTransitionType_HOT_RESTART
	case "HOLD":
		return ModeTransitionType_HOLD
	case "COLD_RESTART":
		return ModeTransitionType_COLD_RESTART
	case "RUN_R":
		return ModeTransitionType_RUN_R
	case "LINK_UP":
		return ModeTransitionType_LINK_UP
	case "UPDATE":
		return ModeTransitionType_UPDATE
	}
	return 0
}

func CastModeTransitionType(structType interface{}) ModeTransitionType {
	castFunc := func(typ interface{}) ModeTransitionType {
		if sModeTransitionType, ok := typ.(ModeTransitionType); ok {
			return sModeTransitionType
		}
		return 0
	}
	return castFunc(structType)
}

func (m ModeTransitionType) LengthInBits() uint16 {
	return 8
}

func (m ModeTransitionType) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func ModeTransitionTypeParse(readBuffer utils.ReadBuffer) (ModeTransitionType, error) {
	val, err := readBuffer.ReadUint8("ModeTransitionType", 8)
	if err != nil {
		return 0, nil
	}
	return ModeTransitionTypeByValue(val), nil
}

func (e ModeTransitionType) Serialize(writeBuffer utils.WriteBuffer) error {
	return writeBuffer.WriteUint8("ModeTransitionType", 8, uint8(e), utils.WithAdditionalStringRepresentation(e.name()))
}

func (e ModeTransitionType) name() string {
	switch e {
	case ModeTransitionType_STOP:
		return "STOP"
	case ModeTransitionType_WARM_RESTART:
		return "WARM_RESTART"
	case ModeTransitionType_RUN:
		return "RUN"
	case ModeTransitionType_HOT_RESTART:
		return "HOT_RESTART"
	case ModeTransitionType_HOLD:
		return "HOLD"
	case ModeTransitionType_COLD_RESTART:
		return "COLD_RESTART"
	case ModeTransitionType_RUN_R:
		return "RUN_R"
	case ModeTransitionType_LINK_UP:
		return "LINK_UP"
	case ModeTransitionType_UPDATE:
		return "UPDATE"
	}
	return ""
}

func (e ModeTransitionType) String() string {
	return e.name()
}
