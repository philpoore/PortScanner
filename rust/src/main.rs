use tokio::net::TcpStream;
use std::error::Error;

use std::env;
use log::{info, error};
use simple_logger::SimpleLogger;

#[derive(Debug)]
struct Stats {
    total: u64,
    open: u64,
    closed: u64,
}

impl Stats {
    fn new() -> Self {
        Stats {total: 0, open: 0, closed: 0}
    }
}

#[tokio::main]
async fn main() -> Result<(), Box<dyn Error>> {
    SimpleLogger::new().with_level(log::LevelFilter::Info).init().unwrap();

    let mut stats = Stats::new();
    
    let host = env::args().nth(1).unwrap_or("127.0.0.1".to_string());
    let max_port_range: i32  = 65535;
    for port in 0..max_port_range {
        let addr = format!("{}:{}", host, port);
        let res = TcpStream::connect(addr.clone()).await;
        stats.total += 1;

        if res.is_err() {
            let err = res.err().unwrap();
            if err.kind() == std::io::ErrorKind::ConnectionRefused {
                stats.closed += 1;
                continue;
            }
            error!("Error: {}", err);
            break;
        }
        stats.open += 1;
        
        info!("Connected {}", addr);
    }

    info!("Results {:?}", stats);
    Ok(())
}
