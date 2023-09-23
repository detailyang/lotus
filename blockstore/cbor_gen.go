// Code generated by github.com/whyrusleeping/cbor-gen. DO NOT EDIT.

package blockstore

import (
	"fmt"
	"io"
	"math"
	"sort"

	cid "github.com/ipfs/go-cid"
	cbg "github.com/whyrusleeping/cbor-gen"
	xerrors "golang.org/x/xerrors"
)

var _ = xerrors.Errorf
var _ = cid.Undef
var _ = math.E
var _ = sort.Sort

var lengthBufNetRpcReq = []byte{132}

func (t *NetRpcReq) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}

	cw := cbg.NewCborWriter(w)

	if _, err := cw.Write(lengthBufNetRpcReq); err != nil {
		return err
	}

	// t.Type (blockstore.NetRPCReqType) (uint8)
	if err := cw.WriteMajorTypeHeader(cbg.MajUnsignedInt, uint64(t.Type)); err != nil {
		return err
	}

	// t.ID (uint64) (uint64)

	if err := cw.WriteMajorTypeHeader(cbg.MajUnsignedInt, uint64(t.ID)); err != nil {
		return err
	}

	// t.Cid ([]cid.Cid) (slice)
	if len(t.Cid) > cbg.MaxLength {
		return xerrors.Errorf("Slice value in field t.Cid was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajArray, uint64(len(t.Cid))); err != nil {
		return err
	}
	for _, v := range t.Cid {

		if err := cbg.WriteCid(cw, v); err != nil {
			return xerrors.Errorf("failed to write cid field v: %w", err)
		}

	}

	// t.Data ([][]uint8) (slice)
	if len(t.Data) > cbg.MaxLength {
		return xerrors.Errorf("Slice value in field t.Data was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajArray, uint64(len(t.Data))); err != nil {
		return err
	}
	for _, v := range t.Data {
		if len(v) > cbg.ByteArrayMaxLen {
			return xerrors.Errorf("Byte array in field v was too long")
		}

		if err := cw.WriteMajorTypeHeader(cbg.MajByteString, uint64(len(v))); err != nil {
			return err
		}

		if _, err := cw.Write(v[:]); err != nil {
			return err
		}
	}
	return nil
}

func (t *NetRpcReq) UnmarshalCBOR(r io.Reader) (err error) {
	*t = NetRpcReq{}

	cr := cbg.NewCborReader(r)

	maj, extra, err := cr.ReadHeader()
	if err != nil {
		return err
	}
	defer func() {
		if err == io.EOF {
			err = io.ErrUnexpectedEOF
		}
	}()

	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 4 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.Type (blockstore.NetRPCReqType) (uint8)

	maj, extra, err = cr.ReadHeader()
	if err != nil {
		return err
	}
	if maj != cbg.MajUnsignedInt {
		return fmt.Errorf("wrong type for uint8 field")
	}
	if extra > math.MaxUint8 {
		return fmt.Errorf("integer in input was too large for uint8 field")
	}
	t.Type = NetRPCReqType(extra)
	// t.ID (uint64) (uint64)

	{

		maj, extra, err = cr.ReadHeader()
		if err != nil {
			return err
		}
		if maj != cbg.MajUnsignedInt {
			return fmt.Errorf("wrong type for uint64 field")
		}
		t.ID = uint64(extra)

	}
	// t.Cid ([]cid.Cid) (slice)

	maj, extra, err = cr.ReadHeader()
	if err != nil {
		return err
	}

	if extra > cbg.MaxLength {
		return fmt.Errorf("t.Cid: array too large (%d)", extra)
	}

	if maj != cbg.MajArray {
		return fmt.Errorf("expected cbor array")
	}

	if extra > 0 {
		t.Cid = make([]cid.Cid, extra)
	}

	for i := 0; i < int(extra); i++ {
		{
			var maj byte
			var extra uint64
			var err error
			_ = maj
			_ = extra
			_ = err

			{

				c, err := cbg.ReadCid(cr)
				if err != nil {
					return xerrors.Errorf("failed to read cid field t.Cid[i]: %w", err)
				}

				t.Cid[i] = c

			}
		}
	}

	// t.Data ([][]uint8) (slice)

	maj, extra, err = cr.ReadHeader()
	if err != nil {
		return err
	}

	if extra > cbg.MaxLength {
		return fmt.Errorf("t.Data: array too large (%d)", extra)
	}

	if maj != cbg.MajArray {
		return fmt.Errorf("expected cbor array")
	}

	if extra > 0 {
		t.Data = make([][]uint8, extra)
	}

	for i := 0; i < int(extra); i++ {
		{
			var maj byte
			var extra uint64
			var err error
			_ = maj
			_ = extra
			_ = err

			maj, extra, err = cr.ReadHeader()
			if err != nil {
				return err
			}

			if extra > cbg.ByteArrayMaxLen {
				return fmt.Errorf("t.Data[i]: byte array too large (%d)", extra)
			}
			if maj != cbg.MajByteString {
				return fmt.Errorf("expected byte array")
			}

			if extra > 0 {
				t.Data[i] = make([]uint8, extra)
			}

			if _, err := io.ReadFull(cr, t.Data[i][:]); err != nil {
				return err
			}
		}
	}

	return nil
}

var lengthBufNetRpcResp = []byte{131}

func (t *NetRpcResp) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}

	cw := cbg.NewCborWriter(w)

	if _, err := cw.Write(lengthBufNetRpcResp); err != nil {
		return err
	}

	// t.Type (blockstore.NetRPCRespType) (uint8)
	if err := cw.WriteMajorTypeHeader(cbg.MajUnsignedInt, uint64(t.Type)); err != nil {
		return err
	}

	// t.ID (uint64) (uint64)

	if err := cw.WriteMajorTypeHeader(cbg.MajUnsignedInt, uint64(t.ID)); err != nil {
		return err
	}

	// t.Data ([]uint8) (slice)
	if len(t.Data) > cbg.ByteArrayMaxLen {
		return xerrors.Errorf("Byte array in field t.Data was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajByteString, uint64(len(t.Data))); err != nil {
		return err
	}

	if _, err := cw.Write(t.Data[:]); err != nil {
		return err
	}
	return nil
}

func (t *NetRpcResp) UnmarshalCBOR(r io.Reader) (err error) {
	*t = NetRpcResp{}

	cr := cbg.NewCborReader(r)

	maj, extra, err := cr.ReadHeader()
	if err != nil {
		return err
	}
	defer func() {
		if err == io.EOF {
			err = io.ErrUnexpectedEOF
		}
	}()

	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 3 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.Type (blockstore.NetRPCRespType) (uint8)

	maj, extra, err = cr.ReadHeader()
	if err != nil {
		return err
	}
	if maj != cbg.MajUnsignedInt {
		return fmt.Errorf("wrong type for uint8 field")
	}
	if extra > math.MaxUint8 {
		return fmt.Errorf("integer in input was too large for uint8 field")
	}
	t.Type = NetRPCRespType(extra)
	// t.ID (uint64) (uint64)

	{

		maj, extra, err = cr.ReadHeader()
		if err != nil {
			return err
		}
		if maj != cbg.MajUnsignedInt {
			return fmt.Errorf("wrong type for uint64 field")
		}
		t.ID = uint64(extra)

	}
	// t.Data ([]uint8) (slice)

	maj, extra, err = cr.ReadHeader()
	if err != nil {
		return err
	}

	if extra > cbg.ByteArrayMaxLen {
		return fmt.Errorf("t.Data: byte array too large (%d)", extra)
	}
	if maj != cbg.MajByteString {
		return fmt.Errorf("expected byte array")
	}

	if extra > 0 {
		t.Data = make([]uint8, extra)
	}

	if _, err := io.ReadFull(cr, t.Data[:]); err != nil {
		return err
	}
	return nil
}

var lengthBufNetRpcErr = []byte{131}

func (t *NetRpcErr) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}

	cw := cbg.NewCborWriter(w)

	if _, err := cw.Write(lengthBufNetRpcErr); err != nil {
		return err
	}

	// t.Type (blockstore.NetRPCErrType) (uint8)
	if err := cw.WriteMajorTypeHeader(cbg.MajUnsignedInt, uint64(t.Type)); err != nil {
		return err
	}

	// t.Msg (string) (string)
	if len(t.Msg) > cbg.MaxLength {
		return xerrors.Errorf("Value in field t.Msg was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len(t.Msg))); err != nil {
		return err
	}
	if _, err := cw.WriteString(string(t.Msg)); err != nil {
		return err
	}

	// t.Cid (cid.Cid) (struct)

	if t.Cid == nil {
		if _, err := cw.Write(cbg.CborNull); err != nil {
			return err
		}
	} else {
		if err := cbg.WriteCid(cw, *t.Cid); err != nil {
			return xerrors.Errorf("failed to write cid field t.Cid: %w", err)
		}
	}

	return nil
}

func (t *NetRpcErr) UnmarshalCBOR(r io.Reader) (err error) {
	*t = NetRpcErr{}

	cr := cbg.NewCborReader(r)

	maj, extra, err := cr.ReadHeader()
	if err != nil {
		return err
	}
	defer func() {
		if err == io.EOF {
			err = io.ErrUnexpectedEOF
		}
	}()

	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 3 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.Type (blockstore.NetRPCErrType) (uint8)

	maj, extra, err = cr.ReadHeader()
	if err != nil {
		return err
	}
	if maj != cbg.MajUnsignedInt {
		return fmt.Errorf("wrong type for uint8 field")
	}
	if extra > math.MaxUint8 {
		return fmt.Errorf("integer in input was too large for uint8 field")
	}
	t.Type = NetRPCErrType(extra)
	// t.Msg (string) (string)

	{
		sval, err := cbg.ReadString(cr)
		if err != nil {
			return err
		}

		t.Msg = string(sval)
	}
	// t.Cid (cid.Cid) (struct)

	{

		b, err := cr.ReadByte()
		if err != nil {
			return err
		}
		if b != cbg.CborNull[0] {
			if err := cr.UnreadByte(); err != nil {
				return err
			}

			c, err := cbg.ReadCid(cr)
			if err != nil {
				return xerrors.Errorf("failed to read cid field t.Cid: %w", err)
			}

			t.Cid = &c
		}

	}
	return nil
}
