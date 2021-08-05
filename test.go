package main

//import (
//	"fmt"
//
//	"github.com/bits-and-blooms/bitset"
//)
//
//const DefaultSize = 2 << 24
//
//var seeds = []uint{7, 11, 13, 31, 37, 61}
//
////一个布隆过滤器包含一个二进制数组（用第三方库Bitset），还有哈希函数
//type BloomFilter struct {
//	set       *bitset.BitSet
//	hashFuncs [6]SimpleHash
//}
//
//func NewBloomFilter() *BloomFilter {
//	//造一个布隆过滤器的指针
//	bf := new(BloomFilter)
//	//初始化哈希函数
//	for i := 0; i < len(bf.hashFuncs); i++ {
//		bf.hashFuncs[i] = SimpleHash{DefaultSize, seeds[i]}
//	}
//	//初始化哈希数组
//	bf.set = bitset.New(DefaultSize)
//	return bf
//}
//
//func (b *BloomFilter) add(value string) {
//	if value == "" {
//		return
//	}
//	for _, f := range b.hashFuncs {
//		//先对value做哈希得到索引值
//		//Set函数把对应索引位置设为1
//		b.set.Set(f.hash(value))
//	}
//}
//
//func (b *BloomFilter) contains(value string) bool {
//	if value == "" {
//		return false
//	}
//	result := true
//	for _, f := range b.hashFuncs {
//		result = result && b.set.Test(f.hash(value))
//	}
//	return result
//}
//
////用种子数来区分不同的哈希函数
//type SimpleHash struct {
//	cap  uint
//	seed uint
//}
//
////字符串的哈希函数
//func (s *SimpleHash) hash(value string) uint {
//	var result uint = 0
//	//取字符串中每一个字符，累加到result
//	for i := 0; i < len(value); i++ {
//		result = result*s.seed + uint(value[i])
//	}
//	//和哈希数组容量相与
//	return (s.cap - 1) & result
//}
//
//func main() {
//	filter := NewBloomFilter()
//	filter.add("123")
//	fmt.Println(filter.contains("123"))
//	fmt.Println(filter.contains("456"))
//}
