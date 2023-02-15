# K8s Pod Info API
A simple JSON API that returns some information about the Kubernetes Pod where it's running.

I wrote this to help me debug kubernetes services and load balancers.

## Pod template config
In order to return the pod info. The app expects certain environment variables to be set. You can set them in the pod template this way:
````yaml
apiVersion: v1
kind: Pod
metadata:
...
spec:
...
  containers:
    - ...
      env:
      - name: POD_NAME
          valueFrom:
            fieldRef:
              fieldPath: metadata.name
      - name: POD_IP
          valueFrom:
            fieldRef:
              fieldPath: status.podIP
      - name: POD_NAMESPACE
          valueFrom:
            fieldRef:
              fieldPath: metadata.namespace
      - name: POD_SERVICE_ACCOUNT_NAME
          valueFrom:
            fieldRef:
              fieldPath: spec.serviceAccountName
      - name: NODE_NAME
          valueFrom:
            fieldRef:
              fieldPath: spec.nodeName
````


## Install tools
````bash
make install-tools
````

## Run tests
````bash
make test
````

## Run locally
Run the dev server:
````bash
POD_NAME="my-pod-name" POD_IP="192.168.13.101" POD_NAMESPACE="my-namespace" POD_SERVICE_ACCOUNT_NAME="my-service-account" NODE_NAME="my-node-name" PORT="3000" make run-dev
````

Make a curl request:
````bash
curl -s localhost:3000/api/v1/info
````


## License

MIT License

Copyright (c) 2023 Adil H

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.