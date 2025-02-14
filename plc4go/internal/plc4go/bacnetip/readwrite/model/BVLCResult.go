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
type BVLCResult struct {
	*BVLC
	Code BVLCResultCode
}

// The corresponding interface
type IBVLCResult interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *BVLCResult) BvlcFunction() uint8 {
	return 0x00
}

func (m *BVLCResult) InitializeParent(parent *BVLC, bvlcPayloadLength uint16) {}

func NewBVLCResult(code BVLCResultCode, bvlcPayloadLength uint16) *BVLC {
	child := &BVLCResult{
		Code: code,
		BVLC: NewBVLC(bvlcPayloadLength),
	}
	child.Child = child
	return child.BVLC
}

func CastBVLCResult(structType interface{}) *BVLCResult {
	castFunc := func(typ interface{}) *BVLCResult {
		if casted, ok := typ.(BVLCResult); ok {
			return &casted
		}
		if casted, ok := typ.(*BVLCResult); ok {
			return casted
		}
		if casted, ok := typ.(BVLC); ok {
			return CastBVLCResult(casted.Child)
		}
		if casted, ok := typ.(*BVLC); ok {
			return CastBVLCResult(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *BVLCResult) GetTypeName() string {
	return "BVLCResult"
}

func (m *BVLCResult) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *BVLCResult) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.ParentLengthInBits())

	// Simple field (code)
	lengthInBits += 16

	return lengthInBits
}

func (m *BVLCResult) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func BVLCResultParse(readBuffer utils.ReadBuffer) (*BVLC, error) {
	if pullErr := readBuffer.PullContext("BVLCResult"); pullErr != nil {
		return nil, pullErr
	}

	// Simple Field (code)
	if pullErr := readBuffer.PullContext("code"); pullErr != nil {
		return nil, pullErr
	}
	_code, _codeErr := BVLCResultCodeParse(readBuffer)
	if _codeErr != nil {
		return nil, errors.Wrap(_codeErr, "Error parsing 'code' field")
	}
	code := _code
	if closeErr := readBuffer.CloseContext("code"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BVLCResult"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BVLCResult{
		Code: code,
		BVLC: &BVLC{},
	}
	_child.BVLC.Child = _child
	return _child.BVLC, nil
}

func (m *BVLCResult) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BVLCResult"); pushErr != nil {
			return pushErr
		}

		// Simple Field (code)
		if pushErr := writeBuffer.PushContext("code"); pushErr != nil {
			return pushErr
		}
		_codeErr := m.Code.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("code"); popErr != nil {
			return popErr
		}
		if _codeErr != nil {
			return errors.Wrap(_codeErr, "Error serializing 'code' field")
		}

		if popErr := writeBuffer.PopContext("BVLCResult"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BVLCResult) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
