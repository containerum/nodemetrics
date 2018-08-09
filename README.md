# nodeMetrics
nodeMetrics is a service for [Containerum](https://github.com/containerum/containerum) that monitors node resources utilisation (RAM, CPU, storage). 

## Prerequisites
* Kubernetes

## Installation

### Using Helm

```
  helm repo add containerum https://charts.containerum.io
  helm repo update
  helm install containerum/nodemetrics
```
## Contributions
Please submit all contributions concerning nodeMetrics component to this repository. Contributing guidelines are available [here](https://github.com/containerum/containerum/blob/master/CONTRIBUTING.md).

## License
nodeMetrics project is licensed under the terms of the Apache License Version 2.0. Please see License in this repository for more details.
