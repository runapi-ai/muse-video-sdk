<p align="center">
  <a href="https://github.com/runapi-ai/muse-video">
    <h3 align="center">Muse Video API Skill for RunAPI</h3>
  </a>
</p>

<p align="center">
  Install this agent skill to create Muse Video tasks through RunAPI.
</p>

<div align="center">

[![skills.sh](https://www.skills.sh/b/runapi-ai/muse-video)](https://www.skills.sh/runapi-ai/muse-video/muse-video)
[![ClawHub](https://img.shields.io/badge/ClawHub-runapi--muse--video-111827)](https://clawhub.ai/runapi-ai/runapi-muse-video)
[![License](https://img.shields.io/github/license/runapi-ai/muse-video)](https://github.com/runapi-ai/muse-video/blob/main/LICENSE)

</div>
<br/>

```bash
npx skills add runapi-ai/muse-video -g
```

Or paste this prompt to your AI agent:

```text
Install the muse-video skill for me:

1. Clone https://github.com/runapi-ai/muse-video
2. Copy the skills/muse-video/ directory into your
   user-level skills directory (e.g. ~/.claude/skills/
   for Claude Code, ~/.codex/skills/ for Codex, or the
   configured Gemini CLI skills directory).
3. Verify that SKILL.md is present.
4. Confirm the install path when done.
```

The canonical agent instructions are in `skills/muse-video/SKILL.md` and cover Claude Code, Codex, Gemini CLI, Cursor, and compatible skill hosts.

## Quick Example

```shell
runapi muse-video text-to-video --input-file request.json
```

Use the target-language SDK for application integration. Use the CLI for one-off jobs and manual testing.

## Links

- [Model reference](https://runapi.ai/models/muse-video)
- [API reference](https://runapi.ai/docs/api/muse-video/text-to-video)
- [SDK repository](https://github.com/runapi-ai/muse-video-sdk)
- [Pricing and rate limits](https://runapi.ai/models/muse-video)

## License

Licensed under the Apache License, Version 2.0.
