![CI][ci-status]
[![PkgGoDev][pkg-go-dev-badge]][pkg-go-dev]

# go-aws-cloudwatch-event-details

go-aws-cloudwatch-event-details provides more `detail` field types of CloudWatch Events.

## Usage

```go
import (
  "context"

  "github.com/aereal/go-aws-cloudwatch-event-details"
  "github.com/aws/aws-lambda-go/events"
)

func handler(ctx context.Context, event events.CloudWatchEvent) error {
  detail, err := cweventdetails.ParseEventDetail(event)
  if err != nil{
    return err
  }
  println(detail.(*cweventdetails.ECSTaskStateChangeEvent).StoppedReason)
  return nil
}
```

## Install

```sh
go get github.com/aereal/go-aws-cloudwatch-event-details
```

## License

See LICENSE file.

[pkg-go-dev]: https://pkg.go.dev/github.com/aereal/go-aws-cloudwatch-event-details
[pkg-go-dev-badge]: https://pkg.go.dev/badge/aereal/go-aws-cloudwatch-event-details
[ci-status]: https://github.com/aereal/go-aws-cloudwatch-event-details/workflows/CI/badge.svg?branch=main
