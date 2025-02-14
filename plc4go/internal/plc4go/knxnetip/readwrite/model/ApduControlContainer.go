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

// The data-structure of this message
type ApduControlContainer struct {
	*Apdu
	ControlApdu *ApduControl
}

// The corresponding interface
type IApduControlContainer interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *ApduControlContainer) Control() uint8 {
	return uint8(1)
}

func (m *ApduControlContainer) InitializeParent(parent *Apdu, numbered bool, counter uint8) {
	m.Numbered = numbered
	m.Counter = counter
}

func NewApduControlContainer(controlApdu *ApduControl, numbered bool, counter uint8) *Apdu {
	child := &ApduControlContainer{
		ControlApdu: controlApdu,
		Apdu:        NewApdu(numbered, counter),
	}
	child.Child = child
	return child.Apdu
}

func CastApduControlContainer(structType interface{}) *ApduControlContainer {
	castFunc := func(typ interface{}) *ApduControlContainer {
		if casted, ok := typ.(ApduControlContainer); ok {
			return &casted
		}
		if casted, ok := typ.(*ApduControlContainer); ok {
			return casted
		}
		if casted, ok := typ.(Apdu); ok {
			return CastApduControlContainer(casted.Child)
		}
		if casted, ok := typ.(*Apdu); ok {
			return CastApduControlContainer(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *ApduControlContainer) GetTypeName() string {
	return "ApduControlContainer"
}

func (m *ApduControlContainer) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *ApduControlContainer) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.ParentLengthInBits())

	// Simple field (controlApdu)
	lengthInBits += m.ControlApdu.LengthInBits()

	return lengthInBits
}

func (m *ApduControlContainer) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func ApduControlContainerParse(readBuffer utils.ReadBuffer, dataLength uint8) (*Apdu, error) {
	if pullErr := readBuffer.PullContext("ApduControlContainer"); pullErr != nil {
		return nil, pullErr
	}

	// Simple Field (controlApdu)
	if pullErr := readBuffer.PullContext("controlApdu"); pullErr != nil {
		return nil, pullErr
	}
	_controlApdu, _controlApduErr := ApduControlParse(readBuffer)
	if _controlApduErr != nil {
		return nil, errors.Wrap(_controlApduErr, "Error parsing 'controlApdu' field")
	}
	controlApdu := CastApduControl(_controlApdu)
	if closeErr := readBuffer.CloseContext("controlApdu"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("ApduControlContainer"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &ApduControlContainer{
		ControlApdu: CastApduControl(controlApdu),
		Apdu:        &Apdu{},
	}
	_child.Apdu.Child = _child
	return _child.Apdu, nil
}

func (m *ApduControlContainer) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ApduControlContainer"); pushErr != nil {
			return pushErr
		}

		// Simple Field (controlApdu)
		if pushErr := writeBuffer.PushContext("controlApdu"); pushErr != nil {
			return pushErr
		}
		_controlApduErr := m.ControlApdu.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("controlApdu"); popErr != nil {
			return popErr
		}
		if _controlApduErr != nil {
			return errors.Wrap(_controlApduErr, "Error serializing 'controlApdu' field")
		}

		if popErr := writeBuffer.PopContext("ApduControlContainer"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *ApduControlContainer) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
