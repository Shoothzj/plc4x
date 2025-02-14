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
type BVLCRegisterForeignDevice struct {
	*BVLC
	Ttl uint16
}

// The corresponding interface
type IBVLCRegisterForeignDevice interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *BVLCRegisterForeignDevice) BvlcFunction() uint8 {
	return 0x05
}

func (m *BVLCRegisterForeignDevice) InitializeParent(parent *BVLC, bvlcPayloadLength uint16) {}

func NewBVLCRegisterForeignDevice(ttl uint16, bvlcPayloadLength uint16) *BVLC {
	child := &BVLCRegisterForeignDevice{
		Ttl:  ttl,
		BVLC: NewBVLC(bvlcPayloadLength),
	}
	child.Child = child
	return child.BVLC
}

func CastBVLCRegisterForeignDevice(structType interface{}) *BVLCRegisterForeignDevice {
	castFunc := func(typ interface{}) *BVLCRegisterForeignDevice {
		if casted, ok := typ.(BVLCRegisterForeignDevice); ok {
			return &casted
		}
		if casted, ok := typ.(*BVLCRegisterForeignDevice); ok {
			return casted
		}
		if casted, ok := typ.(BVLC); ok {
			return CastBVLCRegisterForeignDevice(casted.Child)
		}
		if casted, ok := typ.(*BVLC); ok {
			return CastBVLCRegisterForeignDevice(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *BVLCRegisterForeignDevice) GetTypeName() string {
	return "BVLCRegisterForeignDevice"
}

func (m *BVLCRegisterForeignDevice) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *BVLCRegisterForeignDevice) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.ParentLengthInBits())

	// Simple field (ttl)
	lengthInBits += 16

	return lengthInBits
}

func (m *BVLCRegisterForeignDevice) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func BVLCRegisterForeignDeviceParse(readBuffer utils.ReadBuffer) (*BVLC, error) {
	if pullErr := readBuffer.PullContext("BVLCRegisterForeignDevice"); pullErr != nil {
		return nil, pullErr
	}

	// Simple Field (ttl)
	_ttl, _ttlErr := readBuffer.ReadUint16("ttl", 16)
	if _ttlErr != nil {
		return nil, errors.Wrap(_ttlErr, "Error parsing 'ttl' field")
	}
	ttl := _ttl

	if closeErr := readBuffer.CloseContext("BVLCRegisterForeignDevice"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BVLCRegisterForeignDevice{
		Ttl:  ttl,
		BVLC: &BVLC{},
	}
	_child.BVLC.Child = _child
	return _child.BVLC, nil
}

func (m *BVLCRegisterForeignDevice) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BVLCRegisterForeignDevice"); pushErr != nil {
			return pushErr
		}

		// Simple Field (ttl)
		ttl := uint16(m.Ttl)
		_ttlErr := writeBuffer.WriteUint16("ttl", 16, (ttl))
		if _ttlErr != nil {
			return errors.Wrap(_ttlErr, "Error serializing 'ttl' field")
		}

		if popErr := writeBuffer.PopContext("BVLCRegisterForeignDevice"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BVLCRegisterForeignDevice) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
