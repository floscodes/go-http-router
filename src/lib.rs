// imports for hyper
// and make them publicly accessible in this crate
pub use hyper::*;

// use standard core Result, not the one from hyper itself
pub use core::result::Result;

mod error_fns;
mod route_config;
mod static_route_config;
mod serve;

use route_config::*;
use static_route_config::*;

#[derive(Clone)]
pub struct Router {
    routes: Vec<RouteConfig>,
    static_routes: Vec<StaticRouteConfig>,
}

pub fn new_router() -> RouterBuilder {
    RouterBuilder {
        router: Router{
        routes: vec![],
        static_routes: vec![],
    },
    }
}


pub struct RouterBuilder {
    router: Router,
}

impl RouterBuilder {

pub fn finish(self) -> Router {
    self.router
}

pub fn handle(mut self, route: &'static str, handlerfn: fn(req: Request<Body>) -> Response<Body>) -> Self {
    let rc = RouteConfig {
        route: route,
        handlerfn: handlerfn,
        methods: vec![],
        accept_trailing_slash: true,
        allow_cors: false,
    };
    self.router.routes.push(rc);
    self
}

pub fn allow_method(mut self, method: &'static str) -> Self {
    let l = self.router.routes.len();
    if l > 1 {
        self.router.routes[l-1].methods.push(method);
    }
    self
}

pub fn allow_cors(mut self, b: bool) -> Self {
    let l = self.router.routes.len();
    if l > 1 {
        self.router.routes[l-1].allow_cors=b;
    }
    self
}

pub fn allow_trailing_slash(mut self, b: bool) -> Self {
    let l = self.router.routes.len();
    if l > 1 {
        self.router.routes[l-1].accept_trailing_slash=b;
    }
    self
}

pub fn serve_static(mut self, urlpath: &'static str, dirpath: &'static str) -> Self {
    let rc = StaticRouteConfig {
        route: urlpath,
        dirpath: dirpath,
        index_file: "",
    };
    self.router.static_routes.push(rc);
    self
}

}

    // Start server
    pub async fn start_server(addr: &'static str, router: Router) -> Result<(),()> {
        serve::serve(addr, router).await
    }


// Tests
async fn test_methods() {
    let rt= new_router()
        .handle("/", |_req| {Response::new("Test".into())}).finish();

    start_server(":8080", rt).await.unwrap();

}
