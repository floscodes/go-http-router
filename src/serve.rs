// imports for hyper
use std::convert::Infallible;
use std::net::SocketAddr;
use crate::error_fns::*;
use crate::*;
use hyper::{Body, Server};
use hyper::service::{make_service_fn, service_fn};
use std::str::FromStr;

// Start server
pub (in super) async fn serve(addr: &'static str, router: Router) -> Result<(),()> {
    let addr = match SocketAddr::from_str(addr) {
        Ok(a) => a,
        Err(_) => return Err(())
    };

    let make_svc = make_service_fn(move |_conn| {
        let router_capture1 = router.clone();
        async move {
        
        Ok::<_, Infallible>(service_fn(move |req: Request<Body>| {
            
            let rt = router_capture1.clone();

            async move {
            
            // HANDLE THE CONFIGURED ROUTES
            for mut route in rt.routes {
                //remove trailing slash, if one exists
                route.route = route.route.trim_end_matches("/");
                // handle trailing slash settings with if clause
                if route.accept_trailing_slash && format!("{}/", route.route) == req.uri().to_string() || route.route == &req.uri().to_string() {

                    // Other checks
                    if !route.check_methods(&req) {
                        return Ok(forbidden()) as Result<Response<Body>, Infallible>
                    }

                    // TODO HERE: GET CHECK ALLOW CORS! (Write function before!!)

                }
            }
        
            Ok(notfound()) as Result<Response<Body>, Infallible>
        }}))
}
});

    let server = Server::bind(&addr).serve(make_svc);

    if let Err(_e) = server.await {
        return Err(());
    }

    Ok(())
}
