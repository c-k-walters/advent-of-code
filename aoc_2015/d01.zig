// Advent of Code 2015
// Day 01
// Charlie Kelley

const std = @import("std");
const input = @embedFile("./inputs/d01.txt");
const out = std.debug;

pub fn main() !void {
    out.print("==========   AoC 2015 Day 01   ==========\n", .{});

    var floor: i32 = 0;
    var firstBase: i32 = 0;
    var found: bool = false;

    var it = std.mem.window(u8, input, 1, 1);
    var i: i32 = 1;
    while (it.next()) |token| : (i += 1) {
        switch (token[0]) {
            '(' => {
                floor += 1;
            },
            ')' => {
                floor -= 1;
            },
            else => {},
        }
        if (!found and floor < 0) {
            firstBase = i;
            found = true;
        }
    }

    out.print("Part 1 Solution: Floor {}\n", .{floor});
    out.print("Part 2 Solution: Floor {}\n", .{firstBase});

    out.print("==========================================\n", .{});
}
