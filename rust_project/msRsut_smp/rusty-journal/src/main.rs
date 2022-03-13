mod cli;
use structopt::StructOpt;
mod tasks;

use cli::{Action::*, CommandLineArgs};
use tasks::Task;

use std::path::PathBuf;
use anyhow::anyhow;

fn find_default_journal_file() -> Option<PathBuf> {
    home::home_dir().map(|mut path| {
        path.push(".rusty-journal.json");
        path
    })
}

fn main() {
    let CommandLineArgs {
        action,
        journal_file,
    } = CommandLineArgs::from_args();

    let journal_file = journal_file.or_else(find_default_journal_file).ok_or(anyhow!("Failed to find journal file."))?;

    match action {
        Add { task } => tasks::add_task(journal_file, Task::new(task)),
        List => tasks::liset_tasks(journal_file),
        Done { position } => tasks::complete_task(journal_file, position),
    }?;
    Ok(())
}
