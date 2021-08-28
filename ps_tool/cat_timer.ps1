function clear_con {
  [char]0x1B+"[2J"
  [char]0x1B+"[H"
}

function rep_emoji ($emoji) {
  $StrippedUnicode = $emoji -replace 'U\+',''
  $UnicodeInt = [System.Convert]::toInt32($StrippedUnicode,16)
  $Repemoji = [System.Char]::ConvertFromUtf32($UnicodeInt)

  return $Repemoji 
}

$VALTIME = 40

$PERCENT_GAUGE = "<100%==============75%==================50%=================25%==================0%>"

$CAT_EMOJI = rep_emoji 'U+1F408'
$FOOD_EMOJI = rep_emoji 'U+1F969'
$CONG_EMOJI = rep_emoji 'U+1F389'
$TIMER_EMOJI = rep_emoji 'U+1F55B'

if([string]::IsNullorEmpty($Args[0]))
{
  throw 'No arguments are specified.'
}
elseif([int]$Args[0] -lt 1) 
{
  throw 'Do not specify a number less than or equal to 0.
         Specify a natural number greater than or equal to 1.
         or Do not specify any characters.'
}

$str = '' 
for ($i=0; $i -lt $VALTIME; $i++) {
  $str += '_'
}

# minute exchange
$min_time = [int]$Args[0] * 60

# calc 1%
$one_per_num = $min_time / 100

# guage fit number calc
$guage_num = $one_per_num * 2.5

$cong_str = $CONG_EMOJI + $Args[0] + " minutes passed" + $TIMER_EMOJI + $CONG_EMOJI

$val = $VALTIME
for ($i=0; $i -le $VALTIME; $i++) {
  clear_con
  Write-Host $PERCENT_GAUGE -ForegroundColor DarkYellow
  Write-Host $FOOD_EMOJI $str[0..$val] $CAT_EMOJI -ForegroundColor DarkMagenta 
  $val--
  sleep $guage_num
}
clear_con
Write-Host $FOOD_EMOJI $CAT_EMOJI
Write-Host $cong_str
