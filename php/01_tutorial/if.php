<?php
if (strpos($_SERVER['HTTP_USER_AGENT'], 'MSIE') !== false) {
  echo 'あなたはInternet Explorerを使用しています<br />';
} else {
  echo 'あなたはInternet Explorerを使用していません<br />';
}
?>
