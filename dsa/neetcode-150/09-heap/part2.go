/*
===============================================================
Exercise 6: Design Twitter
===============================================================
Question:
Design simplified Twitter with postTweet, getNewsFeed, follow, unfollow.

Example:
twitter.postTweet(1, 5);
twitter.getNewsFeed(1); // [5]

Key Idea:
HashMap for user data + merge k sorted lists for news feed.

===============================================================
Exercise 7: Find Median from Data Stream
===============================================================
Question:
Design data structure to find median of stream.

Example:
addNum(1);
addNum(2);
findMedian(); // 1.5

Key Idea:
Two heaps: max heap for lower half, min heap for upper half.

===============================================================
*/
package main

import "fmt"

// ===============================================================
// Exercise 6: Design Twitter - Solution
// ===============================================================
type Twitter struct {
	// TODO: Implement
}

func ConstructorTwitter() Twitter {
	return Twitter{}
}

func (this *Twitter) PostTweet(userId int, tweetId int) {
	// TODO: Implement
}

func (this *Twitter) GetNewsFeed(userId int) []int {
	// TODO: Implement
	return nil
}

func (this *Twitter) Follow(followerId int, followeeId int) {
	// TODO: Implement
}

func (this *Twitter) Unfollow(followerId int, followeeId int) {
	// TODO: Implement
}

// Time Complexity: Post O(1), Feed O(n log k), Follow/Unfollow O(1)
// Space Complexity: O(users + tweets)

// ===============================================================
// Exercise 7: Find Median from Data Stream - Solution
// ===============================================================
type MedianFinder struct {
	// TODO: Implement
}

func ConstructorMedian() MedianFinder {
	return MedianFinder{}
}

func (this *MedianFinder) AddNum(num int) {
	// TODO: Implement
}

func (this *MedianFinder) FindMedian() float64 {
	// TODO: Implement
	return 0.0
}

// Time Complexity: AddNum O(log n), FindMedian O(1)
// Space Complexity: O(n)

// ===============================================================
// Main - Test all solutions
// ===============================================================
func main() {
	fmt.Println("===== Heap Part 2 Practice =====")
	fmt.Println("Complete the TODO sections above")
}
