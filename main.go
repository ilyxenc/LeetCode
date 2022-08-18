package main

import (
	"LeetCode/tasks"
	"fmt"
)

func main()  {

	l1 := &tasks.ListNode{
		Val: 1,
		Next: &tasks.ListNode{
			Val: 0,
			Next: &tasks.ListNode{
				Val: 0,
				Next: &tasks.ListNode{
					Val: 0,
					Next: &tasks.ListNode{
						Val: 0,
						Next: &tasks.ListNode{
							Val: 0,
							Next: &tasks.ListNode{
								Val: 0,
								Next: &tasks.ListNode{
									Val: 0,
									Next: &tasks.ListNode{
										Val: 0,
										Next: &tasks.ListNode{
											Val: 0,
											Next: &tasks.ListNode{
												Val: 0,
												Next: &tasks.ListNode{
													Val: 0,
													Next: &tasks.ListNode{
														Val: 0,
														Next: &tasks.ListNode{
															Val: 0,
															Next: &tasks.ListNode{
																Val: 0,
																Next: &tasks.ListNode{
																	Val: 0,
																	Next: &tasks.ListNode{
																		Val: 0,
																		Next: &tasks.ListNode{
																			Val: 0,
																			Next: &tasks.ListNode{
																				Val: 0,
																				Next: &tasks.ListNode{
																					Val: 0,
																					Next: &tasks.ListNode{
																						Val: 1,
																						Next: nil,
																					},
																				},
																			},
																		},
																	},
																},
															},
														},
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	l2 := &tasks.ListNode{
		Val: 5,
		Next: &tasks.ListNode{
			Val: 6,
			Next: &tasks.ListNode{
				Val: 4,
				Next: nil,
			},
		},
	}

	result := tasks.AddTwoNumbers(l1, l2)

	fmt.Println(result)
}