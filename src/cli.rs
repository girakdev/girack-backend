use clap::Clap;

#[derive(Clap)]
#[clap(version = "0.1.0")]
pub struct Opts {
    /// Controll a database for test
    #[clap(short, long, global = true)]
    pub test: bool,

    #[clap(subcommand)]
    pub subcmd: SubCommand,
}

#[derive(Clap)]
pub enum SubCommand {
    Run(Run),
    Db(Db),
}

/// Run server
#[derive(Clap)]
pub struct Run {}

/// A subcommand for controlling a database
#[derive(Clap)]
pub struct Db {
    #[clap(subcommand)]
    pub subcmd: DbSubCommand,
}

#[derive(Clap)]
pub enum DbSubCommand {
    Create(DbCreate),
    Init(DbInit),
    Drop(DbDrop),
}

/// Create a database
#[derive(Clap)]
pub struct DbCreate {}

/// Initialize a database
#[derive(Clap)]
pub struct DbInit {}

/// Drop a database
#[derive(Clap)]
pub struct DbDrop {}

pub fn get_opts() -> Opts {
    Opts::parse()
}

pub fn is_using_test_db() -> bool {
    get_opts().test
}
