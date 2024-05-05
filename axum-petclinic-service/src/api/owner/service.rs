use crate::error::Error;
use crate::repository::owner;
use crate::repository::schema::Owner;

// trait Servicer {
    // fn new(&self, ) -> Self;
    // async fn find_owners(&self) -> Result<Vec<Owner>, Error>;

// }

#[derive(Debug)]
struct Service {
    repo:  owner::Repository,
}

impl Service {
    // fn new(&self) -> Self {
    //     Self {repo: owner::Repository::new()}
    // }
    //
    //
    // async fn find_owners(&self) -> Result<Vec<Owner>, Error> {
    //     self.repo.list_owners()
    // }
    async pub fn new() -> Self {
        Self {
            repo: owner::Repository::new(),
        }
    }
    async fn find_owners(&self) -> Result<Vec<Owner>, Error> {
        self.repo.list_owners()
    }
}
