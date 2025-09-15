use abstract_interface::{Abstract, AccountI};
use abstract_std::objects::gov_type::GovernanceDetails;
use cw_orch_daemon::{networks::BITSONG_2B, RUNTIME};

use clap::Parser;
use cw_orch::prelude::*;
use interchain_bitsong_accounts::{assert_wallet_balance, BITSONG_LOCAL_1, BITSONG_LOCAL_2};

pub const ABSTRACT_VERSION: &str = env!("CARGO_PKG_VERSION");
pub const DEPLOYMENT_DAO: &str =
    "bitsong13hmdq0slwmff7sej79kfa8mgnx4rl46nj2fvmlgu6u32tz6vfqesdfq4vm";

// Run "cargo run --example download_wasms" in the `abstract-interfaces` package before deploying!
fn full_deploy(mut networks: Vec<ChainInfoOwned>, authz_granter: Option<String>) -> anyhow::Result<()> {
    // let networks = RUNTIME.block_on(assert_wallet_balance(networks));

    for network in networks {
        let mut chain = DaemonBuilder::new(network.clone()).build()?;
        
        // Conditionally set authz granter based on environment or parameter
        if let Some(granter) = &authz_granter {
            println!("Using AuthZ granter: {}", granter);
            chain
                .sender_mut()
                .set_authz_granter(&Addr::unchecked(granter.to_string()));
        } else {
            println!("Using direct sender (no AuthZ)");
        }

        let monarch = chain.sender_addr();

        let deployment = match Abstract::deploy_on(chain, ()) {
            Ok(deployment) => {
                // write_deployment(&deployment_status)?;
                deployment
            }
            Err(e) => {
                // write_deployment(&deployment_status)?;
                return Err(e.into());
            }
        };

        // Create the Abstract Account because it's needed for the fees for the dex module
        AccountI::create_default_account(
            &deployment,
            GovernanceDetails::Monarchy {
                monarch: monarch.to_string(),
            },
        )?;
    }

    // fs::copy(Path::new("~/.cw-orchestrator/state.json"), to)
    Ok(())
}

#[derive(Parser, Default, Debug)]
#[command(author, version, about, long_about = None)]
struct Arguments {
    /// Network Id to deploy on
    #[arg(short, long, value_delimiter = ' ', num_args = 1..)]
    network_ids: Vec<String>,
    
    /// AuthZ granter address (optional - if not provided, direct signing is used)
    #[arg(short, long)]
    authz_granter: Option<String>,
}

fn main() {
    // Load environment variables
    dotenv::from_path(".env").ok();
    let mnemonic = dotenv::var("LOCAL_MNEMONIC").unwrap();
    env_logger::init();
    println!("{:#?}", mnemonic);
    
    let args = Arguments::parse();

    // Determine authz usage from environment variable or command line
    let use_authz = dotenv::var("USE_AUTHZ").unwrap_or_else(|_| "false".to_string());
    let authz_granter = if use_authz == "true" {
        args.authz_granter.or_else(|| Some(DEPLOYMENT_DAO.to_string()))
    } else {
        args.authz_granter  // Use CLI arg even if env says false (CLI overrides env)
    };

    let networks = vec![BITSONG_LOCAL_1.into(), BITSONG_LOCAL_2.into()];

    // let networks = args
    //     .network_ids
    //     .iter()
    //     .map(|n| parse_network(n).unwrap().into())
    //     .collect::<Vec<_>>();

    if let Err(ref err) = full_deploy(networks, authz_granter) {
        log::error!("{}", err);
        err.chain()
            .skip(1)
            .for_each(|cause| log::error!("because: {}", cause));

        // The backtrace is not always generated. Try to run this example
        // with `$env:RUST_BACKTRACE=1`.
        //    if let Some(backtrace) = e.backtrace() {
        //        log::debug!("backtrace: {:?}", backtrace);
        //    }

        ::std::process::exit(1);
    }
}
