use crate::*;

pub (in super) fn forbidden() -> Response<Body> {
    Response::builder().status(StatusCode::FORBIDDEN).body(Body::empty()).unwrap()
}

pub (in super) fn notfound() -> Response<Body> {
    Response::builder().status(StatusCode::NOT_FOUND).body(Body::empty()).unwrap()
}