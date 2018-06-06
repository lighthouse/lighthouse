# Lighthouse: A Helm Chart Visualizer

Lighthouse visually displays the template components of Helm charts along with the associated template YAML.

![Image](https://raw.githubusercontent.com/lighthouse/lighthouse/master/images/screenshot.png)

## Usage

Build a docker image or use the stable version at yaron2/lighthouse:

```
docker run -d -p 8000:8000 yaron2/lighthouse
```

Point your browser to localhost:8000.

## Future plans

* Create and show a graph of the different template components
