# openfaas-imagecrawlerdemux

An [openfaas](https://www.openfaas.com/) function which is an image crawler that returns a list of jpg and png images on a site, and sends those lists to [openfaas-exiffeed](https://github.com/servernull/openfaas-exiffeed), [openfaas-nsfwfeed](https://github.com/servernull/openfaas-opennsfwfeed), and [openfaas-inceptionfeed](https://github.com/servernull/openfaas-inceptionfeed).

```bash
# deploy
faas-cli deploy -f stack.yml

# invoke synchronously
echo http://scottleedavis.com | faas-cli invoke openfaas-imagecrawlerdemux | jq

[
  "http://scottleedavis.com/assets/img/achilles.jpg",
  "http://scottleedavis.com/assets/img/ijmuiden-gps-data-1.jpg",
  "http://scottleedavis.com/assets/img/linkedin.png",
  "http://scottleedavis.com/assets/img/scott.jpeg",
  "https://github.com/scottleedavis/google-earth-toolbox/raw/master/screenshot.png",
  "http://scottleedavis.com/assets/img/circadia.png",
  "http://scottleedavis.com/assets/img/light_art.png",
  "http://scottleedavis.com/assets/img/github.png",
  "http://scottleedavis.com/assets/img/twitter.png",
  "http://scottleedavis.com/assets/img/tortoise.jpg"
]

```
