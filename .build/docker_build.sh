# at root dir exec

docker build -f ./Examples/SimpleWeb/Dockerfile  -t liangboceo/yuanboot:v-20201104-56b0d607160cac3954d21f545bcd644541667309 .

kubectl create configmap yuanboot-demo-test -n yuanboot --from-file=config_test.yml