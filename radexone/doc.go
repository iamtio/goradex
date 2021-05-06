/*
Package radexone is reverse engineered serial protocol for Radex One - personal dosimeter from Quarta.
https://quartarad.com/product/radex-one/

Mostly based on https://github.com/mwwhited work:
https://gist.github.com/mwwhited/2491b47c3483f1ac319c227f0dd3a814
and
https://github.com/mwwhited/BinaryDataDecoders/tree/master/src/BinaryDataDecoders.Quarta.RadexOne

Usage

Example
 drr := radexone.NewDataRequest(0)
 encoded := drr.Marshal() // encoded now contains []byte to send
 ...
 resp := radexone.DataReadResponse{}
 resp.Unmarshal(buf) // buf is received []byte
 fmt.Printf("CPM: %d, Ambient: %d, Accumulated: %d", resp.CPM, resp.Ambient, resp.Accumulated)
*/
package radexone
