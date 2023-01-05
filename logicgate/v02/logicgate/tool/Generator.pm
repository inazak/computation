# logicgate Generator.pm
use strict;
use warnings;

my $width = 4;
my $count = 1;
my @buffer = ();
my $FS = sprintf("%%0%dd", $width);

sub LO {
  return "9900";
}

sub HI {
  return "9901";
}

sub CLOCK {
  return "9909";
}

sub SetCounter {
  my ($c) = @_;
  $count = $c;
}

sub FormatString {
  return $FS;
}

sub WriteDefine {
  printf("LO %s\n", LO());
  printf("HI %s\n", HI());
  printf("CLOCK %s\n", CLOCK());
}

sub Write {
  printf("%s\n", join("\n", @buffer) );
  @buffer = ();
}

sub _Nand {
  my $a = $count++;
  my $b = $count++;
  my $o = $count++;

  push(@buffer, sprintf("NAND $FS $FS $FS", $a, $b, $o));

  return ($a, $b, $o);
}

sub Nand {
  my ($x, $y) = @_;

  my ($a, $b, $o) = _Nand();
  Connect($x, $a);
  Connect($y, $b);

  return $o;
}

sub Latch {
  my ($x, $y) = @_;

  my ($a, $v, $q) = _Nand();
  my ($p, $b, $w) = _Nand();
  Connect($q, $p);
  Connect($w, $v);
  Connect($x, $a);
  Connect($y, $b);

  return $q;
}

sub Connect {
  my ($a, $b) = @_;

  push(@buffer, sprintf("CONNECT $FS $FS", $a, $b));
}

sub Not {
  my ($x) = @_;

  my $o = Nand($x, $x);
  
  return $o;
}

sub And {
  my ($x, $y) = @_;

  my $p = Not(Nand($x, $y));
  
  return $p;
}

sub Or {
  my ($x, $y) = @_;

  my $o = Nand(Not($x), Not($y));
  
  return $o;
}

# DLatch
# if enable is HI(1), then q equals d value.
# if enable is LO(0), then q keep prev value.
# 
#  d   e   | q
#  --------|----
#  LO  LO  | q (latched)
#  HI  LO  | q (latched)
#  LO  HI  | LO
#  HI  HI  | HI
sub DLatch {
  my ($d, $e) = @_;

  my $a = Nand($d, $e);
  my $b = Nand(Not($d), $e);
  my $q = Latch($a, $b);

  return $q;
}

sub DFF {
  my ($d, $clock) = @_;

  my $a = DLatch($d, Not($clock));
  my $q = DLatch($a, $clock);

  return $q;
}

# DLatchC
sub DLatchC {
  my ($d, $e, $clear) = @_;

  my $a = Nand($d, $e);
  my $b = Nand(Not($d), $e);
  my $g = Or($a, $clear);
  my $h = And($b, Not($clear));
  my $q = Latch($g, $h);

  return $q;
}

sub DFFC {
  my ($d, $clock, $clear) = @_;

  my $a = DLatchC($d, Not($clock), $clear);
  my $q = DLatchC($a, $clock, $clear);

  return $q;
}

sub And3 {
  my ($a, $b, $c) = @_;
  return And($c, And($a, $b));
}

sub Or3 {
  my ($a, $b, $c) = @_;
  return Or($c, Or($a, $b));
}

sub And4 {
  my ($a, $b, $c, $d) = @_;
  return And($d, And($c, And($a, $b)));
}

sub Or4 {
  my ($a, $b, $c, $d) = @_;
  return Or($d, Or($c, Or($a, $b)));
}

sub Xor {
  my ($a, $b) = @_;
  my $x = Nand($a, $b);
  my $y = Nand($a, $x);
  my $z = Nand($b, $x);
  return Nand($y, $z);
}

# Mux is Multiplexor
# if sel is LO, then output a, otherwise output b.
#
#  sel | out
#  LO  | a
#  HI  | b
#
#  a   b   sel | out
#  ------------|----
#  LO  LO  LO  | LO
#  LO  HI  LO  | LO
#  HI  LO  LO  | HI
#  HI  HI  LO  | HI
#  LO  LO  HI  | LO
#  LO  HI  HI  | HI
#  HI  LO  HI  | LO
#  HI  HI  HI  | HI
#
sub Mux {
  my ($a, $b, $sel) = @_;
  return Or( And($a, Not($sel)), And($b, $sel) );
}

# Mux2 is Multiplexor has 2bit selector
#
#  s0   s1   | out
#  LO   LO   | a
#  HI   LO   | b
#  LO   HI   | c
#  HI   HI   | d
#
sub Mux2 {
  my ($a, $b, $c, $d, $s0, $s1) = @_;
  return Or(And(Mux($a, $b, $s0), Not($s1)), And(Mux($c, $d, $s0), $s1));
}

# BDec is Binary Decoder
#
#   sel | q0  q1
#   LO  | HI  LO
#   HI  | LO  HI
#
sub BDec {
  my ($sel) = @_;

  my $q0 = Not($sel);
  my $q1 = $count++;
  push(@buffer, sprintf("CONNECT $FS $FS", $sel, $q1));

  return $q0, $q1;
}

# BDec2 is Binary Decoder has 2bit selector
#
#  s0  s1 | q0  q1  q2  q3
#  LO  LO | HI  LO  LO  LO
#  HI  LO | LO  HI  LO  LO
#  LO  HI | LO  LO  HI  LO
#  HI  HI | LO  LO  LO  HI
#
sub BDec2 {
  my ($s0, $s1) = @_;

  my $q0 = And(Not($s0), Not($s1));
  my $q1 = And(    $s0 , Not($s1));
  my $q2 = And(Not($s0),     $s1 );
  my $q3 = And(    $s0 ,     $s1 );
  return $q0, $q1, $q2, $q3;
}

# BDec3 is Binary Decoder has 3bit selector
#
#  s0  s1  s2 | q0  q1  q2  q3  q4  q5  q6  q7
#  LO  LO  LO | HI  LO  LO  LO  LO  LO  LO  LO
#  HI  LO  LO | LO  HI  LO  LO  LO  LO  LO  LO
#  LO  HI  LO | LO  LO  HI  LO  LO  LO  LO  LO
#  HI  HI  LO | LO  LO  LO  HI  LO  LO  LO  LO
#  LO  LO  HI | LO  LO  LO  LO  HI  LO  LO  LO
#  HI  LO  HI | LO  LO  LO  LO  LO  HI  LO  LO
#  LO  HI  HI | LO  LO  LO  LO  LO  LO  HI  LO
#  HI  HI  HI | LO  LO  LO  LO  LO  LO  LO  HI
#
sub BDec3 {
  my ($s0, $s1, $s2) = @_;

  my @q = ();
  $q[0] = And3(Not($s0), Not($s1), Not($s2));
  $q[1] = And3(    $s0 , Not($s1), Not($s2));
  $q[2] = And3(Not($s0),     $s1 , Not($s2));
  $q[3] = And3(    $s0 ,     $s1 , Not($s2));
  $q[4] = And3(Not($s0), Not($s1),     $s2 );
  $q[5] = And3(    $s0 , Not($s1),     $s2 );
  $q[6] = And3(Not($s0),     $s1 ,     $s2 );
  $q[7] = And3(    $s0 ,     $s1 ,     $s2 );
  return @q;
}

sub BDec4 {
  my ($s0, $s1, $s2, $s3) = @_;

  my @q = ();
  $q[0]  = And4(Not($s0), Not($s1), Not($s2), Not($s3));
  $q[1]  = And4(    $s0 , Not($s1), Not($s2), Not($s3));
  $q[2]  = And4(Not($s0),     $s1 , Not($s2), Not($s3));
  $q[3]  = And4(    $s0 ,     $s1 , Not($s2), Not($s3));
  $q[4]  = And4(Not($s0), Not($s1),     $s2 , Not($s3));
  $q[5]  = And4(    $s0 , Not($s1),     $s2 , Not($s3));
  $q[6]  = And4(Not($s0),     $s1 ,     $s2 , Not($s3));
  $q[7]  = And4(    $s0 ,     $s1 ,     $s2 , Not($s3));
  $q[8]  = And4(Not($s0), Not($s1), Not($s2),     $s3 );
  $q[9]  = And4(    $s0 , Not($s1), Not($s2),     $s3 );
  $q[10] = And4(Not($s0),     $s1 , Not($s2),     $s3 );
  $q[11] = And4(    $s0 ,     $s1 , Not($s2),     $s3 );
  $q[12] = And4(Not($s0), Not($s1),     $s2 ,     $s3 );
  $q[13] = And4(    $s0 , Not($s1),     $s2 ,     $s3 );
  $q[14] = And4(Not($s0),     $s1 ,     $s2 ,     $s3 );
  $q[15] = And4(    $s0 ,     $s1 ,     $s2 ,     $s3 );
  return @q;
}


# FAdd is 1bit FullAdder
sub Fadd {
  my ($a, $b, $cin) = @_;

  my $w = Nand($a, $b);
  my $x = Nand($b, $cin);
  my $y = Nand($a, $cin);
  my $c = Or(Not($w), Or(Not($x), Not($y)));
  my $s = Xor(Xor($a, $b), $cin);
  return $s, $c;
}

# 3bit FullAdder
sub Fadd3 {
  my ($a0,$a1,$a2, $b0,$b1,$b2, $c0) = @_;

  my ($s0, $c1) = Fadd($a0, $b0, $c0);
  my ($s1, $c2) = Fadd($a1, $b1, $c1);
  my ($s2, $c3) = Fadd($a2, $b2, $c2);

  return $s0,$s1,$s2, $c3;
}


# 4bit FullAdder
sub Fadd4 {
  my ($a0,$a1,$a2,$a3, $b0,$b1,$b2,$b3, $c0) = @_;

  my ($s0, $c1) = Fadd($a0, $b0, $c0);
  my ($s1, $c2) = Fadd($a1, $b1, $c1);
  my ($s2, $c3) = Fadd($a2, $b2, $c2);
  my ($s3, $c4) = Fadd($a3, $b3, $c3);

  return $s0,$s1,$s2,$s3, $c4;
}



# 2bit counter, loadable
sub Counter2 {
  my ($d0, $d1, $clock, $load, $clear) = @_;

  my $r0  = $count++;
  my $q0  = DFFC($r0, $clock, $clear);
  my $fb0 = Not($q0);
  my $in0 = Mux($fb0, $d0, $load);
  push(@buffer, sprintf("CONNECT $FS $FS", $in0, $r0));

  my $r1  = $count++;
  my $q1  = DFFC($r1, $clock, $clear);
  my $fb1 = Xor($q0, $q1);
  my $in1 = Mux($fb1, $d1, $load);
  push(@buffer, sprintf("CONNECT $FS $FS", $in1, $r1));

  return $q0, $q1;
}

# 3bit counter, loadable
sub Counter3 {
  my ($d0, $d1, $d2, $clock, $load, $clear) = @_;

  my $r0  = $count++;
  my $q0  = DFFC($r0, $clock, $clear);
  my $fb0 = Not($q0);
  my $in0 = Mux($fb0, $d0, $load);
  push(@buffer, sprintf("CONNECT $FS $FS", $in0, $r0));

  my $r1  = $count++;
  my $q1  = DFFC($r1, $clock, $clear);
  my $fb1 = Xor($q0, $q1);
  my $in1 = Mux($fb1, $d1, $load);
  push(@buffer, sprintf("CONNECT $FS $FS", $in1, $r1));

  my $r2  = $count++;
  my $q2  = DFFC($r2, $clock, $clear);
  my $a1  = And($q1, Xor($q0, $q2));
  my $a2  = And($q2, Not($q1));
  my $fb2 = Or($a1, $a2);
  my $in2 = Mux($fb2, $d2, $load);
  push(@buffer, sprintf("CONNECT $FS $FS", $in2, $r2));

  return $q0, $q1, $q2;
}

# 4bit counter, loadable
sub Counter4 {
  my ($d0, $d1, $d2, $d3, $clock, $load, $clear) = @_;

  my $r0  = $count++;
  my $q0  = DFFC($r0, $clock, $clear);
  my $fb0 = Not($q0);
  my $in0 = Mux($fb0, $d0, $load);
  push(@buffer, sprintf("CONNECT $FS $FS", $in0, $r0));

  my $r1  = $count++;
  my $q1  = DFFC($r1, $clock, $clear);
  my $fb1 = Xor($q0, $q1);
  my $in1 = Mux($fb1, $d1, $load);
  push(@buffer, sprintf("CONNECT $FS $FS", $in1, $r1));

  my $r2  = $count++;
  my $q2  = DFFC($r2, $clock, $clear);
  my $a1  = And($q1, Xor($q0, $q2));
  my $a2  = And($q2, Not($q1));
  my $fb2 = Or($a1, $a2);
  my $in2 = Mux($fb2, $d2, $load);
  push(@buffer, sprintf("CONNECT $FS $FS", $in2, $r2));

  my $r3  = $count++;
  my $q3  = DFFC($r3, $clock, $clear);
  my $b1  = And(Not($q2), $q3);
  my $b2  = And4($q0, $q1, $q2, Not($q3));
  my $b3  = And3(Not($q0), $q2, $q3);
  my $b4  = And3(Not($q1), $q2, $q3);
  my $fb3 = Or4($b1, $b2, $b3, $b4);
  my $in3 = Mux($fb3, $d3, $load);
  push(@buffer, sprintf("CONNECT $FS $FS", $in3, $r3));

  return $q0, $q1, $q2, $q3;
}


# 3bit Register
sub Register3 {
  my ($d0, $d1, $d2, $clock, $load, $clear) = @_;

  my $r0  = $count++;
  my $q0  = DFFC($r0, $clock, $clear);
  my $in0 = Mux($q0, $d0, $load);
  push(@buffer, sprintf("CONNECT $FS $FS", $in0, $r0));

  my $r1  = $count++;
  my $q1  = DFFC($r1, $clock, $clear);
  my $in1 = Mux($q1, $d1, $load);
  push(@buffer, sprintf("CONNECT $FS $FS", $in1, $r1));

  my $r2  = $count++;
  my $q2  = DFFC($r2, $clock, $clear);
  my $in2 = Mux($q2, $d2, $load);
  push(@buffer, sprintf("CONNECT $FS $FS", $in2, $r2));

  return $q0, $q1, $q2;
}


# 4bit Register
sub Register4 {
  my ($d0, $d1, $d2, $d3, $clock, $load, $clear) = @_;

  my $r0  = $count++;
  my $q0  = DFFC($r0, $clock, $clear);
  my $in0 = Mux($q0, $d0, $load);
  push(@buffer, sprintf("CONNECT $FS $FS", $in0, $r0));

  my $r1  = $count++;
  my $q1  = DFFC($r1, $clock, $clear);
  my $in1 = Mux($q1, $d1, $load);
  push(@buffer, sprintf("CONNECT $FS $FS", $in1, $r1));

  my $r2  = $count++;
  my $q2  = DFFC($r2, $clock, $clear);
  my $in2 = Mux($q2, $d2, $load);
  push(@buffer, sprintf("CONNECT $FS $FS", $in2, $r2));

  my $r3  = $count++;
  my $q3  = DFFC($r3, $clock, $clear);
  my $in3 = Mux($q3, $d3, $load);
  push(@buffer, sprintf("CONNECT $FS $FS", $in3, $r3));

  return $q0, $q1, $q2, $q3;
}




1;

__END__

=head1 NAME

  logicgate Generator

=head1 SYNOPSIS

  use FindBin;
  use lib $FindBin::Bin;

  require Generator;

  SetCounter(1);
  WriteDefine();

  my $f = FormatString();

  my $a = -1;
  my $b = -1;
  my $o = -1;

  $a = LO();
  $b = LO();
  $o = And($a, $b);
  printf("//And($f, $f) -> $f \n", $a, $b, $o);
  Write();

  # result is
  #
  # DEFINE 9900 LO
  # DEFINE 9901 HI
  # DEFINE 9909 CLOCK
  # //And(9900, 9900) -> 0006
  # 0001 0002 NAND 0003
  # 9900 CONNECT 0001
  # 9900 CONNECT 0002
  # 0004 0005 NAND 0006
  # 0003 CONNECT 0004
  # 0003 CONNECT 0005

