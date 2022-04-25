<h1 align="center">
  Products Leaderboard
  <br>
</h1>
<h4 align="center">The concept of a leaderboard—a scoreboard showing the ranked names and current scores (or other data points) of the leading competitors—is essential to the world of computer gaming, but leaderboards are now about more than just games. In this case we are using it to create a ranking of the most liked products</h4>
<p align="center">
  <a href="#key-features">Key Features</a> •
  <a href="#dependencies">Dependencies</a> •
  <a href="#how-to-use">How To Use</a> •
  <a href="#documentation">Documentation</a>
</p>

## Key Features

* Massive scale across millions of products
* Decentralized microservice architecture
* Onion architecture
* Built on top of [Redis Leaderboard](https://redis.com/solutions/use-cases/leaderboards/)
* CI / CD pipeline with GitHub Actions
* K8s deployment

## Dependencies

* [Go (1.17) or later](https://go.dev/)
* [Google pubsub-emulator](https://cloud.google.com/pubsub/docs/emulator)
* [Redis](https://redis.io/docs/getting-started)

## How To Use

### Running the project

Before running the project please ensure that all the dependencies are installed in your system. Then follow the next:

1. Run pubsub emulator. [Click here](https://cloud.google.com/pubsub/docs/emulator)
2. Create topic and subscription.

    - Topic: `voucher-metric`
    - Subscription: `voucher-metric-sub`

3. Run Redis. [Click here](https://redis.io/docs/getting-started/installation)
4. Run the project itself.

    ```
    make web
    ```

### Running the tests

In order to run the project tests you need to execute the following command:

```
make test
```

### TODO:
- [ ] Integration tests
- [ ] Circuit breaker
- [ ] Logging
- [ ] Monitoring
- [ ] Error handling
- [ ] Documentation

## Documentation



* If you want to add new features to this project please [see the contribution guide](.github/CONTRIBUTING.md)
* Questions?, <a href="mailto:machester4@gmail.com?Subject=Question about Project" target="_blank">write here</a>