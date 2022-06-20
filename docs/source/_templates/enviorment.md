
# <Enviorment>

## Enviorment Creation

```

```

## Parameters

| Parameter | Type | Description | Default |
| --------- | ---- | ----------- | ------- |

## Observation Space

The observation space matches the API request from battlesnake, you can find this in their official [API Docs](https://docs.battlesnake.com/references/api#post-move).


## Action Space

The Action space consists of a Move, which has 4 diffrent options:
- `up`
- `down`
- `left`
- `right`

## Rewards

| Win | Draw | Loss |
| --- | ---- | ---- |
| 1   | 0    | -1   |

## Verion History

v0 - Initial Release
