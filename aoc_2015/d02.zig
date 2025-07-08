// Advent of Code 2015
// Day 02
// Charlie Kelley

const std = @import("std");
const input = @embedFile("./inputs/d02.txt");
const out = std.debug;

pub fn main() !void {
    out.print("==========   AoC 2015 Day 02   ==========\n", .{});

    var it = std.mem.tokenizeScalar(u8, input, '\n');
    while (it.next()) |token| {
        _ = token;
    }

    out.print("==========================================\n", .{});
}

/// determine the amount of wrapping paper to cover a right rect prism
/// found from the formula f(l, w, h) = 2*l*w + 2*l*h + 2*w*h + extra
/// where extra is equal to the smallest side's surface area.
fn wrappingForBox(dim: []const u32) u32 {
    const to_sort: []u32 = dim[0..];
    std.mem.sort(u32, to_sort, {}, comptime std.sort.asc(u32));

    const extra: u32 = to_sort[0] * to_sort[1];
    var paper: u32 = 0;

    var it = std.mem.window(u32, to_sort, 2, 1);
    while (it.next()) |dims| {
        paper += dims[0] * dims[1] * 2;
    }

    return extra + paper;
}

test "wrapping calcs" {
    const sample = [3]u32{ 2, 3, 4 };
    try std.testing.expectEqual(58, wrappingForBox(&sample));
}
