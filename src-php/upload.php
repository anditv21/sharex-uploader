<?php

include "functions.php";
include "config.php";

if (isset($_POST["key"])) {
  if (in_array($_POST["key"], $tokens)) {
    $result = GeneratePage($length, $_FILES, generateRandomString($length));

    if ($result) {
      echo 'https://' . $_SERVER['SERVER_NAME'] . "/$mediafolder" . "/$result/";
    } else {
      echo "Couldn't Move File";
    }
  } else {
    echo "Authentication Error.";
  }
} else {
  echo "API Key not set.";
}





?>
