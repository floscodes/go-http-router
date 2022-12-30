#[derive(Clone)]
pub struct StaticRouteConfig {
    pub (in super) route: &'static str,
    pub (in super) dirpath: &'static str,
    pub (in super) index_file: &'static str,
}

impl StaticRouteConfig {
    pub fn index_file(&mut self, file: &'static str) {
        self.index_file=file;
    }
}