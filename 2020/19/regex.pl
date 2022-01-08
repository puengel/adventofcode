#!/usr/bin/perl

my $regex   = $ARGV[0];
my $in   = $ARGV[1];

# print "$regex\n";
# print "$in\n";

# $bar = "This is foo and again foo";

$in =~ m/$regex/;

# print "\n$in\n";
# print "$&\n";

$isEqual = ($& eq $in);

# print "Full: ${$& eq $in}\n";

exit($isEqual);