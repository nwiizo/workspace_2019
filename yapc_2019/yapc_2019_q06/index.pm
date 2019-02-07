sub handle (Str $req) { 
  if "Q" ~~ m/ $req / {
    return "Perl6 if word ~~ / regex /"	  
  }

  my regex reg { <$req> };
  #is OK
  loop (my $i = 0; $i < 1000; $i++) {
    for 'ok.txt'.IO.lines -> $line {
      # Do something with $line
      if $line ~~ / <reg> / {
      } else {
        return "[NG]Good Case:$line ans:$req";
      }
    }
    #is NG
    for 'ng.txt'.IO.lines -> $line {
      # Do something with $line
      if $line ~~ / <reg> / {
        return "[NG]Bad Case:$line ans:$req";
      }
    }
  }
  return "[OK]congratulations. If you are the fastest please call the staff. ans:$req";
}

my $st = prompt "";
my Bool $boolean;

$boolean = ?$st;
if !$boolean {
  say "Please challenge the problem! Hope you enjoy it.";
}

my $ret = handle($st);
say "$ret";
