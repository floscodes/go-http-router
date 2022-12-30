use crate::*;
#[derive(Clone)]
pub struct RouteConfig {
    pub (in super) route: &'static str,
    pub (in super) handlerfn: fn(req: Request<Body>) -> Response<Body>,
    pub (in super) methods: Vec<&'static str>,
    pub (in super) accept_trailing_slash: bool,
    pub (in super) allow_cors: bool,
}

impl RouteConfig {
    pub (in crate) fn check_methods(&self, req: &Request<Body>) -> bool {
        if self.methods.len() == 0 {
            return true
        }
        for m in &self.methods {
            if m.to_lowercase() == req.method().as_str().to_lowercase() {
                return true
            }
        }
        return false
    }
}
