use anyhow::Context;

fn run_instructions(content: &str) -> anyhow::Result<()> {
    for inst in content.trim().split("\n") {
        let (cmd, val) = inst.split_at(1);
        let cmd = cmd.chars().nth(0).context("cmd incorrect")?;
        let val: i32 = val.parse()?;
        dbg!(cmd);
        dbg!(val);
    }

    return Ok(());
}

fn main() -> anyhow::Result<()> {
    let _ = run_instructions("");
    return Ok(());
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn run_instructions_test() {
        let sample = "F10
N3
F7
R90
F11";
        let _ = run_instructions(sample);
    }
}
