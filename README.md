## gospider

[gospider](https://github.com/cnych/gospider) is a tiny command-line utility to download media contents (videos, audios, images) from the Web, in case there is no other handy way to do it.

### Download a video

Here's how you use `gospider` to download a video from [秒拍](http://m.miaopai.com/show/channel/aOJK9yvtR-HvLTuhVEeLaVsmFBRfEKLo8i3TaA__):

```console
$ gospider http://m.miaopai.com/show/channel/aOJK9yvtR-HvLTuhVEeLaVsmFBRfEKLo8i3TaA__
site:                           Miaopai
title:                          林更新视频混剪—追光者.mp4
size:                           6.79 MiB (7116238 bytes)
Downloading 林更新视频混剪—追光者.mp4 ...
 6.79 MiB / 6.79 MiB [=====================================================================] 100.00% 7s
Saving Me at the 林更新视频混剪—追光者.mp4 ...Done.
```

### Set the path and name of downloaded file

Use the --output-dir/-o option to set the path, and --output-filename/-O to set the name of the downloaded file:

```console
$ gospider -o ~/Videos -O test.mp4 http://m.miaopai.com/show/channel/aOJK9yvtR-HvLTuhVEeLaVsmFBRfEKLo8i3TaA__
```

#### Tips:

* These options are helpful if you encounter problems with the default video titles, which may contain special characters that do not play well with your current shell / operating system / filesystem.
* These options are also helpful if you write a script to batch download files and put them into designated folders with designated names.