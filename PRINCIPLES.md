# Principles

 * Simplicity of design.
 * The code is the spec.
 * Minimal code - keep total footprint small.
 * Minimal dependencies - all dependencies must get audited, and become part of the repo.
 * Modular dependencies - whereever reasonable, make components modular.
 * Finished - software projects that don't become finished are projects that
   are forever vulnerable. One of the primary goals of the Gno language and
   related works is to become finished within a reasonable timeframe.

## Tendermint Principles

 * Each node can run on a commodity machine. Corollarily, for scaling we focus on sharding & forms of IBC.

## Cli Principles

 * No envs.
 * No short flags.
 * No /bin/ calls.
 * No process forks.
 * Struct-based command options.

## Token Principles

 * Single base token.
 * Deflationary is sufficient when tx fees are imminent.
 * Int64 is sufficiently large to handle human numbers; for everything else, use denom conversions.
