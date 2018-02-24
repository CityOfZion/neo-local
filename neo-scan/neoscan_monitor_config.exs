use Mix.Config

config :neoscan_monitor, ecto_repos: []

config :neoscan_monitor, seeds: [
      System.get_env("NEO_SEED_1"),
      System.get_env("NEO_SEED_2"),
      System.get_env("NEO_SEED_3"),
      System.get_env("NEO_SEED_4")
    ]