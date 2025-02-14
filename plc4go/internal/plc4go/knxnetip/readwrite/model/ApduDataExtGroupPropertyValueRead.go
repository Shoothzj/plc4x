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
)

// Code generated by code-generation. DO NOT EDIT.

// The data-structure of this message
type ApduDataExtGroupPropertyValueRead struct {
	*ApduDataExt
}

// The corresponding interface
type IApduDataExtGroupPropertyValueRead interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *ApduDataExtGroupPropertyValueRead) ExtApciType() uint8 {
	return 0x28
}

func (m *ApduDataExtGroupPropertyValueRead) InitializeParent(parent *ApduDataExt) {}

func NewApduDataExtGroupPropertyValueRead() *ApduDataExt {
	child := &ApduDataExtGroupPropertyValueRead{
		ApduDataExt: NewApduDataExt(),
	}
	child.Child = child
	return child.ApduDataExt
}

func CastApduDataExtGroupPropertyValueRead(structType interface{}) *ApduDataExtGroupPropertyValueRead {
	castFunc := func(typ interface{}) *ApduDataExtGroupPropertyValueRead {
		if casted, ok := typ.(ApduDataExtGroupPropertyValueRead); ok {
			return &casted
		}
		if casted, ok := typ.(*ApduDataExtGroupPropertyValueRead); ok {
			return casted
		}
		if casted, ok := typ.(ApduDataExt); ok {
			return CastApduDataExtGroupPropertyValueRead(casted.Child)
		}
		if casted, ok := typ.(*ApduDataExt); ok {
			return CastApduDataExtGroupPropertyValueRead(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *ApduDataExtGroupPropertyValueRead) GetTypeName() string {
	return "ApduDataExtGroupPropertyValueRead"
}

func (m *ApduDataExtGroupPropertyValueRead) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *ApduDataExtGroupPropertyValueRead) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.ParentLengthInBits())

	return lengthInBits
}

func (m *ApduDataExtGroupPropertyValueRead) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func ApduDataExtGroupPropertyValueReadParse(readBuffer utils.ReadBuffer, length uint8) (*ApduDataExt, error) {
	if pullErr := readBuffer.PullContext("ApduDataExtGroupPropertyValueRead"); pullErr != nil {
		return nil, pullErr
	}

	if closeErr := readBuffer.CloseContext("ApduDataExtGroupPropertyValueRead"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &ApduDataExtGroupPropertyValueRead{
		ApduDataExt: &ApduDataExt{},
	}
	_child.ApduDataExt.Child = _child
	return _child.ApduDataExt, nil
}

func (m *ApduDataExtGroupPropertyValueRead) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ApduDataExtGroupPropertyValueRead"); pushErr != nil {
			return pushErr
		}

		if popErr := writeBuffer.PopContext("ApduDataExtGroupPropertyValueRead"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *ApduDataExtGroupPropertyValueRead) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
