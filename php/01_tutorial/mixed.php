<?php
if (strpos($_SERVER['HTTP_USER_AGENT'], 'MSIE') !== false) {
?>
<h3>strposが非falseを返しました</h3>
<center><b>あなたはInternet Explorerを使用しています</b></center>
<?php
} else {
?>
<h3>strposがfalseを返しました</h3>
<center><b>あなたはInternet Explorerを使用していません</b></center>
<?php
}
?>
