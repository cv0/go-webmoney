// Copyright (c) 2016 The Constantin Karataev. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file or at
// https://developers.google.com/open-source/licenses/bsd.
//
// Interface X7. Verifying client's handwritten signature - owner
// of WM Keeper WinPro (Classic).
// This interface is designed for verifying authenticity
// of the handwritten signature in messages, generated by WM Keeper WinPro (Classic) keys. All parameters are passed in Win-1251 encoding.

// https://wiki.wmtransfer.com/projects/webmoney/wiki/Interface_X7

package webmoney

import (
	"encoding/xml"
)

type TestSignRequest struct {
	XMLName xml.Name `xml:"testsign"`
	Wmid    string   `xml:"wmid"`
	Plan    string   `xml:"plan"`
	Sign    string   `xml:"sign"`
}

type TestSignResponse struct {
	XMLName xml.Name `xml:"testsign"`
	Res     string   `xml:"res"`
}

func (w *WmClient) TestSign(t TestSignRequest) (TestSignResponse, error) {
	w.X = X7
	if w.IsClassic() {
		w.Sign = w.Wmid + t.Wmid + t.Plan + t.Sign
	}
	w.Request = t
	result := TestSignResponse{}
	err := w.getResult(&result)
	return result, err
}