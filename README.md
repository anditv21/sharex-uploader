<h1 align="center"> simple-sharex-uploader </h1>
<p align="center"><img src="https://i.ibb.co/8P7NCFQ/image.png"></p>
<details>
<summary>Setup</summary>
<ol>
  <li>Extract the files and upload them to your host.</li>
  <li>Customize the settings in the config.json file according to your preferences.</li>
  <li>Ensure that you set a secure value for the "main" key in the config.json file.</li>
  <li>Add your domain and your upload key to example.sxcu<li>
  <li>Open the example.sxcu file and modify the domain in the "RequestURL" section to match your own domain.</li>
  <li>Open the main window of sharex</li>
  <li>Click on Destinations -> Custom Uploader Settings...</li>
  <li>Click on Import -> From file.. and select the example.sxcu </li>
  <li>To run the ShareX Go Server as a system service:</li>
    <ul>
      <li>Copy the provided systemd service file content to your clipboard.</li>
      <li>Open a terminal and type: <code>sudo nano /etc/systemd/system/sharex-go.service</code></li>
      <li>Paste the contents into the nano text editor.</li>
      <li>Customize the executable path and working directory in the systemd service file. Modify the <code>ExecStart</code> and <code>WorkingDirectory</code> directives accordingly.</li>
      <li>Save the file by pressing <code>Ctrl + O</code>, then press <code>Enter</code> to confirm. Exit nano by pressing <code>Ctrl + X</code>.</li>
      <li>Reload systemd to load the new service: <code>sudo systemctl daemon-reload</code></li>
      <li>Enable the service to start on boot: <code>sudo systemctl enable sharex-go</code></li>
      <li>Start the service: <code>sudo systemctl start sharex-go</code></li>
    </ul>
</ol>
</details>


<details>
<summary>Click here to see a Example of a Discord Image Embed</summary>
<img src="https://i.ibb.co/zH21Jsp/image.png">
</details>
<details>
<summary>Click here to see a Example of a Discord Video Embed</summary>
<img src="https://i.ibb.co/Xs0SkmC/image.png">
</details>
